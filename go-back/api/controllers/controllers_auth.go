package controllers

import (
	"project-hackathon/core/common"

	"github.com/gin-gonic/gin"
)

type ControllerAuth struct {
	service common.ServiceAuth
}

func NewControllerAuth(service common.ServiceAuth) common.ControllerAuth {
	return &ControllerAuth{service}
}

func (c *ControllerAuth) Register(ctx *gin.Context) {

}

func (c *ControllerAuth) Login(ctx *gin.Context) {

}
