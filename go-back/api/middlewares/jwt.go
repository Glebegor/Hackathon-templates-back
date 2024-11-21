package middlewares

import (
	"project-hackathon/core/responses"
	"project-hackathon/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(401, responses.ErrorResponse{Message: "Authorization header is required", Status: 401})
			ctx.Abort()
			return
		}
		authSplit := strings.Split(authHeader, "Bearer ")
		if len(authSplit) != 2 {
			ctx.JSON(401, responses.ErrorResponse{Message: "Invalid authorization header", Status: 401})
			ctx.Abort()
			return
		}
		token := authSplit[1]
		authorized, err := utils.IsAuthorized(token, secret)
		if err != nil {
			ctx.JSON(401, responses.ErrorResponse{Message: err.Error(), Status: 401})
			ctx.Abort()
			return
		}
		if !authorized {
			ctx.JSON(401, responses.ErrorResponse{Message: "Unauthorized", Status: 401})
			ctx.Abort()
			return
		}
		claimsData, err := utils.GetClaimsData(token, secret)
		if err != nil {
			ctx.JSON(401, responses.ErrorResponse{Message: err.Error(), Status: 500})
			ctx.Abort()
			return
		}
		ctx.Set("userId", claimsData.UserId)
		ctx.Set("userName", claimsData.UserName)
		ctx.Set("userEmail", claimsData.Email)

		ctx.Next()
	}
}
