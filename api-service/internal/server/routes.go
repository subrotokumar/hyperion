package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := s.server
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(httprate.LimitByIP(100, 1*time.Minute))

	r.Get("/health", s.healthHandler)

	//* Auth Routes
	r.Get("/auth/{provider}/callback", s.getAuthProviderCallback)
	r.Get("/logout/{provider}", s.logoutProvider)
	r.Get("/auth/{provider}", s.beginAuthProviderCallback)
	r.Get("/auth/providers", s.getProviderList)

	//* Project
	r.With(AuthMiddleware).Route("/project", func(r chi.Router) {
		r.Get("/{projectId}", s.projectDetail)
		r.Post("", s.deployProject)
		// r.Patch("", s.deploy)
		// r.Delete("")
	})

	return r
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"message": "Server is running 🚀",
	}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}
	_, _ = w.Write(jsonResp)
}
