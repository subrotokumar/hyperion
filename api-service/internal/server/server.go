package server

import (
	"context"
	"encoding/json"
	"fmt"
	db "hyperion/internal/db/sql"
	"hyperion/internal/utility/env"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

var print = fmt.Println

type (
	Server struct {
		port      int
		server    *chi.Mux
		store     *db.SQLStore
		validator *validator.Validate
	}
)

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	chi := chi.NewRouter()

	conn, err := pgxpool.New(context.Background(), env.Env.DatabaseUrl)
	if err != nil {
		log.Fatalf("DB Error: %s", err.Error())
	}
	store := db.NewSQLStore(conn)
	NewServer := &Server{
		port:      port,
		server:    chi,
		validator: validator.New(),
		store:     store,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func (s *Server) ApiError(w http.ResponseWriter, status int, message string) {
	apiError := map[string]any{
		"status":  status,
		"massage": message,
	}

	jsonResp, err := json.Marshal(apiError)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}
	_, _ = w.Write(jsonResp)
}
