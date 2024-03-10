package model

import (
	db "hyperion/internal/db/sql"
	"time"
)

type TokenResponse struct {
	RefreshToken Token `json:"refresh_token"`
	AccessToken  Token `json:"access_token"`
}

type Token struct {
	Token  string    `json:"token"`
	Expiry time.Time `json:"expiry"`
}

type ProjectDetailResponse struct {
	Project    db.Project    `json:"project"`
	Deployment db.Deployment `json:"deployment"`
}
