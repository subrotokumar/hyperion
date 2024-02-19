package jwt

import (
	"hyperion/internal/utility/env"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestJwt(t *testing.T) {
	err := env.LoadEnv("../../../.env")
	require.NoError(t, err)
	require.NotEmpty(t, env.Env.AccessTokenSecret)
	payload := &JWTPayload{
		Id:       237283,
		Role:     "user",
		IssuedAt: time.Now(),
	}
	require.NotEmpty(t, payload)

	// TESTING Access TOKEN
	jwt, err := payload.SignAccessToken()

	require.NoError(t, err)
	require.NotEmpty(t, jwt)

	JwtClaim, err := ValidateAccessToken(jwt)
	require.NotEmpty(t, jwt)
	require.NoError(t, err)

	issue1 := payload.IssuedAt.Unix()
	issue2 := JwtClaim.IssuedAt.Unix()
	require.Equal(t, issue1, issue2)

	// TESTING REFRESH TOKEN
	jwt, err = payload.SignRefreshToken()
	require.NotEmpty(t, jwt)
	require.NoError(t, err)

	JwtClaim, err = ValidateRefreshToken(jwt)
	require.NotEmpty(t, jwt)
	require.NoError(t, err)
	require.Equal(t, payload.Id, JwtClaim.Id)
	require.Equal(t, payload.Id, JwtClaim.Id)
	require.Equal(t, payload.Role, JwtClaim.Role)

	issue1 = payload.IssuedAt.Unix()
	issue2 = JwtClaim.IssuedAt.Unix()
	require.Equal(t, issue1, issue2)

	random := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1"

	JwtClaim1, err := ValidateAccessToken(random)
	require.Error(t, err)
	require.Empty(t, JwtClaim1)

	JwtClaim2, err := ValidateRefreshToken(random)
	require.Error(t, err)
	require.Empty(t, JwtClaim2)

}
