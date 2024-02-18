package server

import (
	"encoding/json"
	"net/http"
)

type DeployRequest struct {
	GithubUrl string `json:"githubUrl" validate:"required"`
}

func (s *Server) deploy(w http.ResponseWriter, r *http.Request) {
	var body DeployRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		s.ApiError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	if err := s.validator.Struct(body); err != nil {
		s.ApiError(w, http.StatusBadRequest, "Bad Request")
		return
	}
}
