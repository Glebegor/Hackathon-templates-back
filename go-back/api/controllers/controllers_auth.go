package controllers

import (
	"project-hackathon/core/common"
	"project-hackathon/core/responses"

	"github.com/gin-gonic/gin"
)

type ControllerAuth struct {
	service common.ServiceAuth
}

func NewControllerAuth(service common.ServiceAuth) common.ControllerAuth {
	return &ControllerAuth{service}
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

	token, err := c.service.Login(&input)
	if err != nil {
		ctx.JSON(400, responses.ErrorResponse{Message: err.Error(), Status: 400})
		return
	}

	ctx.JSON(200, responses.SuccessResponse{Message: "User logged in successfully", Status: 200, Data: map[string]interface{}{"token": token}})
}
