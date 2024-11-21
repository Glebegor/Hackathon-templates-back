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

// Register
//	@Summary		Register
//	@Description	Register to account
//	@Tags			Auth v2
//	@Accept			json
//	@Produce		json
//	@Param			input	body		common.UserRegister	true	"User Register"
//	@Success		200		{object}	responses.SuccessResponse
//	@Failure		400		{object}	responses.ErrorResponse
//	@Failure		500		{object}	responses.ErrorResponse
//	@Router			/api/v2/auth/register [post]
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

// Login
//	@Summary		Login
//	@Description	Login to account
//	@Tags			Auth v2
//	@Accept			json
//	@Produce		json
//	@Param			input	body		common.UserLogin	true	"User Login"
//	@Success		200		{object}	responses.SuccessResponse
//	@Failure		400		{object}	responses.ErrorResponse
//	@Failure		500		{object}	responses.ErrorResponse
//	@Router			/api/v2/auth/login [post]
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

// Refresh
//	@Summary		Refresh
//	@Description	Refresh of the token
//	@Tags			Auth v2
//	@Accept			json
//	@Produce		json
//	@Param			input	body		common.Refresh	true	"User refresh token"
//	@Success		200		{object}	responses.SuccessResponse
//	@Failure		400		{object}	responses.ErrorResponse
//	@Failure		500		{object}	responses.ErrorResponse
//	@Router			/api/v2/auth/refresh [post]
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
