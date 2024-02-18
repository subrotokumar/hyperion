package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	printf  = fmt.Printf
	print   = fmt.Print
	println = fmt.Println
	panic   = log.Panic
)

func GetFileContentType(ouput *os.File) (string, error) {
	buf := make([]byte, 512)
	_, err := ouput.Read(buf)
	if err != nil {
		return "", err
	}
	contentType := http.DetectContentType(buf)
	return contentType, nil
}
