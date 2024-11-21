package middlewares

import (
	"errors"
	"os"
	"project-hackathon/core/common"
	"project-hackathon/core/responses"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWTIndentification(c *gin.Context) {
	header := c.Request.Header.Get("Authorization")
	if header == "" {
		c.JSON(401, responses.ErrorResponse{Message: "Authorization header is required", Status: 401})
		c.Abort()
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.JSON(401, responses.ErrorResponse{Message: "Authorization header must have 2 parts", Status: 401})
		c.Abort()
		return
	}

	var jwtClaims common.JWTClaims
	token := headerParts[1]
	jwtClaims, err := ParseJWT(token)
	if err != nil {
		c.JSON(401, responses.ErrorResponse{Message: err.Error(), Status: 401})
		c.Abort()
		return
	}
	c.Set("userName", jwtClaims.Name)
	c.Set("userEmail", jwtClaims.Email)
	c.Set("userId", jwtClaims.Id)
	c.Next()
}

func ParseJWT(token string) (common.JWTClaims, error) {
	tokenAccess, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid signing method")
		}
		return []byte(os.Getenv("SERVER_SECRET")), nil
	})

	if err != nil {
		return common.JWTClaims{}, err
	}

	claims, ok := tokenAccess.Claims.(jwt.MapClaims)
	if !ok || !tokenAccess.Valid {
		return common.JWTClaims{}, errors.New("Invalid token")
	}

	jwtClaims := common.JWTClaims{
		Email: claims["email"].(string),
		Name:  claims["name"].(string),
		Id:    int(claims["id"].(float64)),
		Exp:   int64(claims["exp"].(float64)),
	}

	return jwtClaims, nil
}
