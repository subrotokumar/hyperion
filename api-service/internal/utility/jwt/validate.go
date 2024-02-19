package jwt

import (
	"fmt"
	"hyperion/internal/utility/env"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func validateJwt(tokenString string, isAccess bool) (*JWTPayload, error) {
	var hamcSecret string
	if isAccess {
		hamcSecret = env.Env.AccessTokenSecret
	} else {
		hamcSecret = env.Env.RefreshTokenSecret
	}
	fmt.Println(hamcSecret)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(hamcSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		expiry, err := claims.GetExpirationTime()
		if err != nil {
			return nil, fmt.Errorf("invalid Expiration format")
		} else if time.Now().Unix() > expiry.Time.Unix() {
			return nil, fmt.Errorf("token is Expired")
		}

		audience, err := claims.GetAudience()
		if err != nil {
			return nil, fmt.Errorf("invalid token format")
		}

		issuedAt, err := claims.GetIssuedAt()
		if err != nil {
			return nil, fmt.Errorf("invalid token format")
		}

		subject, err := claims.GetSubject()
		if err != nil {
			return nil, fmt.Errorf("invalid Subject Value")
		}
		id, err := strconv.Atoi(subject)
		if err != nil {
			return nil, fmt.Errorf("invalid Subject Value")
		}

		return &JWTPayload{
			Id:       int32(id),
			IssuedAt: issuedAt.Time,
			Role:     audience[0],
		}, nil
	} else {
		return nil, err
	}
}

func ValidateAccessToken(tokenString string) (*JWTPayload, error) {
	return validateJwt(tokenString, true)
}

func ValidateRefreshToken(tokenString string) (*JWTPayload, error) {
	return validateJwt(tokenString, false)
}
