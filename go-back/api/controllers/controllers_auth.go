package controllers

import (
	"project-hackathon/bootstrap"
	"project-hackathon/core/common"
	"project-hackathon/core/responses"
	"project-hackathon/utils"

	"github.com/gin-gonic/gin"
)

type ControllerAuth struct {
	service common.ServiceAuth
	env     *bootstrap.Env
}

func NewControllerAuth(service common.ServiceAuth, env *bootstrap.Env) common.ControllerAuth {
	return &ControllerAuth{service, env}
}

func (c *ControllerAuth) Register(ctx *gin.Context) {
	var input common.UserRegister

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(400, responses.ErrorResponse{Message: err.Error(), Status: 400})
		return
	}

	if err := c.service.Register(&input); err != nil {
		ctx.JSON(400, responses.ErrorResponse{Message: err.Error(), Status: 400})
		return
	}

	ctx.JSON(200, responses.SuccessResponse{Message: "User registered successfully", Status: 200})
	return
}

func (c *ControllerAuth) Login(ctx *gin.Context) {
	var input common.UserLogin

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(400, responses.ErrorResponse{Message: err.Error(), Status: 400})
		return
	}

	password_hash := utils.HashPassword(input.Password, c.env.SERVER_SECRET)

	user, err := c.service.CheckUserByEmailAndPassword(&common.UserLogin{Email: input.Email, Password: password_hash})
	if err != nil {
		ctx.JSON(400, responses.ErrorResponse{Message: err.Error(), Status: 400})
		return
	}
	user.Password = utils.HashPassword(user.Password, c.env.SERVER_SECRET)
	accessToken, err := utils.CreateAccessToken(user, c.env.SERVER_SECRET)
	refreshToken, err := utils.CreateRefreshToken(user, c.env.SERVER_SECRET)

	ctx.JSON(200, responses.SuccessResponse{Message: "User logged in successfully", Status: 200, Data: gin.H{"access_token": accessToken, "refresh_token": refreshToken}})
	return
}

func (c *ControllerAuth) Refresh(ctx *gin.Context) {
	var input common.Refresh

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(400, responses.ErrorResponse{Message: err.Error(), Status: 400})
		return
	}

	id, err := utils.GetIdFromRefreshToken(input.RefreshToken, c.env.SERVER_SECRET)
	if err != nil {
		ctx.JSON(400, responses.ErrorResponse{Message: err.Error(), Status: 400})
		return
	}

	user, err := c.service.GetUserById(id)
	if err != nil {
		ctx.JSON(400, responses.ErrorResponse{Message: err.Error(), Status: 400})
		return
	}
	accessToken, err := utils.CreateAccessToken(user, c.env.SERVER_SECRET)
	if err != nil {
		ctx.JSON(400, responses.ErrorResponse{Message: err.Error(), Status: 400})
		return
	}
	refreshToken, err := utils.CreateRefreshToken(user, c.env.SERVER_SECRET)
	if err != nil {
		ctx.JSON(400, responses.ErrorResponse{Message: err.Error(), Status: 400})
		return
	}

	ctx.JSON(200, responses.SuccessResponse{Message: "Token refreshed successfully", Status: 200, Data: gin.H{"access_token": accessToken, "refresh_token": refreshToken}})
}
