package model

import "time"

type TokenResponse struct {
	RefreshToken Token `json:"refresh_token"`
	AccessToken  Token `json:"access_token"`
}

type Token struct {
	Token  string    `json:"token"`
	Expiry time.Time `json:"expiry"`
}
