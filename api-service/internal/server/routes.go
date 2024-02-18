package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := s.server
	r.Use(middleware.Logger)

	r.Get("/health", s.HelloWorldHandler)

	//* Auth Routes
	r.Get("/auth/{provider}/callback", s.getAuthProviderCallback)
	r.Get("/logout/{provider}", s.logoutProvider)
	r.Get("/auth/{provider}", s.beginAuthProviderCallback)
	r.Get("/auth/providers", s.getProviderList)

	//* Project
	r.Route("/project", func(r chi.Router) {
		r.Get("/deploy", s.deploy)
	})

	//* Deplyments
	r.Route("/project/deployments/", func(r chi.Router) {

	})

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"message": "Server is running 🚀",
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}
	_, _ = w.Write(jsonResp)
}
