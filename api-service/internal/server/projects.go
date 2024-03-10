package server

import (
	"encoding/json"
	"hyperion/internal/aws"
	db "hyperion/internal/db/sql"
	"hyperion/internal/model"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Server) projectDetail(w http.ResponseWriter, r *http.Request) {
	projectId, err := strconv.Atoi(chi.URLParam(r, "projectId"))
	if err != nil {
		s.ApiError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	projectDetail, err := s.store.GetProjectById(r.Context(), int64(projectId))
	if err != nil {
		s.ApiError(w, http.StatusNotFound, "Project not found")
		return
	}
	deployment, err := s.store.GetDeploymentByProjectId(r.Context(), pgtype.Int8{
		Int64: projectDetail.ID,
		Valid: true,
	})
	if err != nil {
		s.ApiError(w, http.StatusNotFound, "Project not found")
		return
	}

	resp := model.ProjectDetailResponse{
		Project:    projectDetail,
		Deployment: deployment,
	}
	if jsonResp, err := json.Marshal(resp); err != nil {
		s.ApiError(w, http.StatusInternalServerError, "Failed to parse json")
		return
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(jsonResp)
	}
}

func (s *Server) deployProject(w http.ResponseWriter, r *http.Request) {
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
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonResp)
}
