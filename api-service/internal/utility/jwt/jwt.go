package jwt

import (
	"time"
)

type TokenType string

const (
	AccessToken  TokenType = "Access Token"
	RefreshToken TokenType = "Refresh Token"
)

const (
	AccessTokenExpiry  = time.Minute * 15
	RefreshTokenExpiry = time.Hour * 24 * 10
)

type JWTPayload struct {
	Id       int32     `json:"sub"`
	IssuedAt time.Time `json:"iat"`
	Role     string    `json:"aud"`
}

func (payload *JWTPayload) AccessTokenExp() time.Time {
	return payload.IssuedAt.Add(AccessTokenExpiry)
}

func (payload *JWTPayload) RefreshTokenExp() time.Time {
	return payload.IssuedAt.Add(RefreshTokenExpiry)
}
