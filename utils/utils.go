package utils

import (
	"errors"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gustapinto/go_jwt/model"
)

var ErrMissingAuthorizationToken = errors.New("missing authorization token")

func ValidateTokenFromHeaders(headers map[string][]string) (string, error) {
	authorization, exists := headers["Authorization"]
	if !exists {
		return "", ErrMissingAuthorizationToken
	}

	token := strings.Split(authorization[0], " ")[1]

	claims := &model.User{}

	jwtKey := []byte(os.Getenv("JWT_KEY"))

	jwtToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return "", err
	}

	if !jwtToken.Valid {
		return "", jwt.ErrSignatureInvalid
	}

	return token, nil
}
