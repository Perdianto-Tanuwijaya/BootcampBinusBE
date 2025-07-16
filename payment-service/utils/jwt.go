package utils

import (
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("rahasia")

func ParseJWT(tokenStr string) (string, error) {
	claims := &jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return "", err
	}

	userID := (*claims)["userId"].(string)
	return userID, nil
}
