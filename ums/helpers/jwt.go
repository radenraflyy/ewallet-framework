package helpers

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ClaimToken struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	jwt.RegisteredClaims
}

var MapTokenType = map[string]time.Duration{
	"token":         time.Hour * 3,
	"refresh_token": time.Hour * 72,
}

var jwtSecret = []byte(GetEnv("APP_SECRET", ""))

func GenerateToken(ctx context.Context, userId uint, username, fullName, tokenType string, now time.Time) (string, error) {

	claimToken := ClaimToken{
		UserID:   userId,
		Username: username,
		FullName: fullName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    GetEnv("APP_NAME", ""),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(MapTokenType[tokenType])),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimToken)
	result, err := token.SignedString(jwtSecret)
	if err != nil {
		return result, fmt.Errorf("failed to generate token: %v", err)
	}

	return result, err
}

func ValidateToken(ctx context.Context, tokenStr string) (*ClaimToken, error) {
	var (
		claimToken *ClaimToken
		ok         bool
	)

	jwtToken, err := jwt.ParseWithClaims(tokenStr, &ClaimToken{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse jwt: %v", err)
	}

	if claimToken, ok = jwtToken.Claims.(*ClaimToken); !ok || !jwtToken.Valid {
		return nil, fmt.Errorf("token invalid")
	}

	return claimToken, nil
}
