package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

const (
	port = ":8000"
)

func main() {
	http.HandleFunc("/", handleRequestAndRedirect)
	http.ListenAndServe(port, nil)
}

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	basePath = os.Getenv("BASE_PATH")
	subdomain := strings.Split(req.Host, ".")[0]
	targetURL := basePath + "/" + subdomain + "/index.html"

	target, err := url.Parse(targetURL)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("subdomain: ", subdomain)
	fmt.Println("Target: ", target)
	res.Header().Set("Content-Type", "text/html")
	reverseProxy := httputil.NewSingleHostReverseProxy(target)
	reverseProxy.ServeHTTP(res, req)
}
