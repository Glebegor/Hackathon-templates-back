package common

import "github.com/gin-gonic/gin"

type ServiceAuth interface {
}

type RepositoryAuth interface {
}

type ControllerAuth interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}
