package grpc_server

import (
	"context"
	db "hyperion/internal/db/sql"
	"hyperion/internal/pb"
	"hyperion/internal/utility/env"
	"log"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	pb.UnimplementedHyperionServer
	port  int
	store *db.SQLStore
}

func NewServer() *Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	conn, err := pgxpool.New(context.Background(), env.Env.DatabaseUrl)
	if err != nil {
		log.Fatalf("DB Error: %s", err.Error())
	}
	store := db.NewSQLStore(conn)
	NewServer := &Server{
		port:  port,
		store: store,
	}
	return NewServer
}
