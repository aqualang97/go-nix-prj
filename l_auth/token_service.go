package main

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

type JWTCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(userID, lifeTimeMinutes int, secret string) (string, error) {
	claims := &JWTCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(lifeTimeMinutes)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func GetTokenFromBearerString(bearerString string) string {
	if bearerString == "" {
		return ""
	}

	parts := strings.Split(bearerString, "Bearer")
	if len(parts) != 2 {
		return ""
	}
	token := strings.TrimSpace(parts[1])

	if len(token) < 1 {
		return ""
	}

	return token
}

func ValidateToken(tokenString, secret string) (*JWTCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTCustomClaims{},
		func(token *jwt.Token) (interface{}, error,
		) {
			return []byte(secret), nil
		})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTCustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("failed to parse token claims")
	}

	return claims, nil
}
