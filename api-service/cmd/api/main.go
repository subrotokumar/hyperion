package main

import (
	"fmt"
	"hyperion/internal/auth"
	"hyperion/internal/grpc_server"
	"hyperion/internal/pb"
	"hyperion/internal/server"
	"hyperion/internal/utility/env"
	"log"
	"net"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := env.LoadEnv(".env"); err != nil {
		log.Fatalf("Env: %v", err.Error())
	}

	auth.NewAuth()
	// runHttpServer()
	runGrpcServer()
}

func runHttpServer() {
	httpServer := server.NewServer()
	err := httpServer.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}

func runGrpcServer() {
	server := grpc_server.NewServer()
	grpcServer := grpc.NewServer()
	pb.RegisterHyperionServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

}
