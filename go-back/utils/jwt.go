package utils

import (
	"fmt"
	"project-hackathon/core/common"
	"project-hackathon/core/domain"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateAccessToken(input domain.User, secret string) (string, error) {
	claims := &common.JwtClaimsAccess{
		UserId:   input.Id,
		UserName: input.Name,
		Email:    input.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(), // 3 hours
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}

func CreateRefreshToken(input domain.User, secret string) (string, error) {
	claims := &common.JwtClaimsRefresh{
		UserId: input.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 24 hours
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	rt, err := token.SignedString([]byte(secret))
	return rt, err
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return false, nil
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetClaimsData(requestToken string, secret string) (*common.JwtData, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid token")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return &common.JwtData{}, err
	}
	claims := token.Claims.(jwt.MapClaims)

	claimsData := &common.JwtData{
		UserId:   claims["user_id"].(string),
		UserName: claims["user_name"].(string),
		Email:    claims["email"].(string),
	}

	return claimsData, nil
}

func GetIdFromRefreshToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid token")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("Invalid token")
	}

	return claims["user_id"].(string), nil
}
