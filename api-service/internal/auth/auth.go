package auth

import (
	"hyperion/internal/utility"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
)

const (
	MaxAge = 86400 * 30
	IsProd = false
)

func NewAuth() {
	store := sessions.NewCookieStore([]byte(utility.Env.SessionKey))
	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProd
	gothic.Store = store
	goth.UseProviders(
		github.New(utility.Env.GithubKey, utility.Env.GithubSecret, "http://localhost:8080/auth/github/callback"),
	)
}
