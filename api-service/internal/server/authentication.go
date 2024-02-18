package server

import (
	"encoding/json"
	"fmt"
	"hyperion/internal/auth"
	db "hyperion/internal/db/sql"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

func (s *Server) getAuthProviderCallback(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	q.Add("provider", chi.URLParam(r, "provider"))
	r.URL.RawQuery = q.Encode()
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Printf("This is the error from completing user auth %v", err)
		return
	}
	print(" User : ", user)

	existingUser, err := s.store.GetUser(r.Context(), int32(1))
	if err != nil && err != pgx.ErrNoRows {
		s.ApiError(w, http.StatusInternalServerError, "DB: Operation Error")
		return
	}
	print(existingUser)

	githubId, err := strconv.Atoi(user.UserID)
	if err != nil {
		s.ApiError(w, http.StatusInternalServerError, "Invalid ID format")
		return
	}
	if err == pgx.ErrNoRows {
		s.store.CreateUser(r.Context(), db.CreateUserParams{
			GithubID: int32(githubId),
			Name:     pgtype.Text{String: user.Name, Valid: true},
			Username: user.NickName,
			Email:    user.Email,
			Avatar:   pgtype.Text{String: user.AvatarURL, Valid: true},
		})
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "jwt_token",
		Path:     "/",
		MaxAge:   auth.MaxAge,
		HttpOnly: true,
		Secure:   auth.IsProd,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "jwt_token",
		Path:     "/",
		MaxAge:   auth.MaxAge,
		HttpOnly: true,
		Secure:   auth.IsProd,
	})

	http.Redirect(w, r, "http://localhost:5173", http.StatusFound)
}

func (s *Server) beginAuthProviderCallback(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	q.Add("provider", chi.URLParam(r, "provider"))
	r.URL.RawQuery = q.Encode()
	gothic.BeginAuthHandler(w, r)
}

func (s *Server) logoutProvider(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	q.Add("provider", chi.URLParam(r, "provider"))
	r.URL.RawQuery = q.Encode()
	gothic.Logout(w, r)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (s *Server) getProviderList(w http.ResponseWriter, r *http.Request) {
	providers := goth.GetProviders()
	var providerList []string = []string{}
	for name, _ := range providers {
		providerList = append(providerList, name)
	}
	jsonRes, err := json.Marshal(map[string]any{
		"providers": providerList,
	})
	if err != nil {
		s.ApiError(w, http.StatusInternalServerError, "Unable to parse Json")
	}
	w.Write(jsonRes)
}
