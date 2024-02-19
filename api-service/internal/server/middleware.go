package server

import (
	"context"
	"encoding/json"
	"hyperion/internal/utility/jwt"
	"log"
	"net/http"
)

type Key string

const USER_KEY Key = "user"

func AuthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		cookieValue, _ := r.Cookie("refresh_token")
		user, err := jwt.ValidateRefreshToken(cookieValue.Value)
		if err != nil {
			apiError := map[string]any{
				"status":  http.StatusUnauthorized,
				"massage": "JWT validation failed",
			}

			jsonResp, err := json.Marshal(apiError)
			if err != nil {
				log.Fatalf("error handling JSON marshal. Err: %v", err)
			}
			_, _ = w.Write(jsonResp)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, USER_KEY, user)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

	}
	return http.HandlerFunc(fn)
}
