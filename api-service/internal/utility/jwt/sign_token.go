package jwt

import (
	"hyperion/internal/utility/env"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (payload *JWTPayload) signToken(isAccess bool) (string, error) {
	var hamcSecret string
	var duration time.Duration
	if isAccess {
		hamcSecret = env.Env.AccessTokenSecret
		duration = AccessTokenExpiry
	} else {
		hamcSecret = env.Env.RefreshTokenSecret
		duration = RefreshTokenExpiry
	}
	claims := jwt.MapClaims{
		"iss": "Hyperion",
		"sub": strconv.Itoa(int(payload.Id)),
		"nbf": payload.IssuedAt.Unix(),
		"iat": payload.IssuedAt.Unix(),
		"exp": payload.IssuedAt.Add(duration).Unix(),
		"aud": "user",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(hamcSecret))
	return tokenString, err
}

func (payload *JWTPayload) SignAccessToken() (string, error) {
	return payload.signToken(true)
}
func (payload *JWTPayload) SignRefreshToken() (string, error) {
	return payload.signToken(false)
}
