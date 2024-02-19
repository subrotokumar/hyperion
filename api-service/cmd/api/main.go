package main

import (
	"fmt"
	"hyperion/internal/auth"
	"hyperion/internal/server"
	"hyperion/internal/utility/env"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if err := env.LoadEnv(".env"); err != nil {
		log.Fatalf("Env: %v", err.Error())
	}

	auth.NewAuth()
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
