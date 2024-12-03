package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWTToken(userID int, tokenSecret string) (string, error) {
	if tokenSecret == "" {
		return "", jwt.ErrInvalidKey
	}

	subject := strconv.Itoa(userID)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "auth-microservice",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		Subject:   subject,
	})

	return token.SignedString([]byte(tokenSecret))
}

func ValidateJWTToken(tokenString, tokenSecret string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)

	if !ok || !token.Valid {
		return 0, jwt.ErrInvalidKey
	}

	userID, err := strconv.Atoi(claims.Subject)

	if err != nil {
		return 0, err
	}

	return userID, nil
}
