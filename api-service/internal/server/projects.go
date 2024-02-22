package server

import (
	"encoding/json"
	"hyperion/internal/aws"
	db "hyperion/internal/db/sql"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Server) deploy(w http.ResponseWriter, r *http.Request) {
	var body struct {
		CreatedBy    pgtype.Int4 `json:"created_by"`
		Name         string      `json:"name"`
		GithubUrl    string      `json:"github_url"`
		Subdomain    pgtype.Text `json:"subdomain"`
		CustomDomain pgtype.Text `json:"custom_domain"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		s.ApiError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	if err := s.validator.Struct(body); err != nil {
		s.ApiError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	project, err := s.store.Queries.CreateProject(r.Context(), db.CreateProjectParams{
		CreatedBy:    body.CreatedBy,
		Name:         body.GithubUrl,
		GithubUrl:    body.GithubUrl,
		Subdomain:    body.CustomDomain,
		CustomDomain: body.CustomDomain,
	})
	if err != nil {
		s.ApiError(w, http.StatusInternalServerError, "DB: Operation error")
		return
	}
	aws.AwsClient.RunTask(r.Context(), body.GithubUrl, body.Subdomain.String)
	jsonResp, err := json.Marshal(project)
	if err != nil {
		s.ApiError(w, http.StatusBadRequest, "Failed to mershal json response")
		return
	}
	time.Sleep(time.Second * 1)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonResp)
}
