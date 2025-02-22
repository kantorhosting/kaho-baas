package utils

import (
	"Kaho_BaaS/internal/apps/account/models"
	"fmt"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var AppName = "Kaho-BaaS"
var JWT_EXPIRATION = 1 * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

func GenerateJWT(secret, userId, email, name string) (string, error) {
	claims := &models.JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    AppName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(JWT_EXPIRATION)),
		},
		UserID: userId,
		Email:  email,
		Name:   name,
	}
	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		slog.Error("generate jwt",
			"err", err,
		)
		return "", fmt.Errorf("Error generating jwt")
	}

	return signedToken, nil
}

func VerifyingJWT(secret, tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signature")
		}

		return []byte(secret), nil
	})

	if err != nil {
		slog.Error("verifying jwt",
			"err", err,
		)
		return nil, fmt.Errorf("Error parsing token")
	}

	if !token.Valid {
		return nil, fmt.Errorf("Token is invalid")
	}

	return token, nil
}
