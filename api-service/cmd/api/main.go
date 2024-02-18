package main

import (
	"fmt"
	"hyperion/internal/auth"
	"hyperion/internal/server"
	"hyperion/internal/utility"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	utility.LoadEnv()
	auth.NewAuth()
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
