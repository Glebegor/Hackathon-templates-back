package utils

import (
	"project-hackathon/core/domain"
	"time"

	"github.com/golang-jwt/jwt"
)

func JWTGeneration(input domain.User, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": input.Email,
		"name":  input.Name,
		"id":    input.Id,
		"exp":   time.Now().Add(time.Hour * 3).Unix(), // 3 hours
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
