package common

import "github.com/golang-jwt/jwt"

type JwtClaimsAccess struct {
	UserId   string `json:"user_id" `
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

type JwtClaimsRefresh struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

type JwtData struct {
	UserId   string `json:"user_id" `
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

type Refresh struct {
	RefreshToken string `json:"refresh_token"`
}
