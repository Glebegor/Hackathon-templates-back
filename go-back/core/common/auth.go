package common

import (
	"github.com/gin-gonic/gin"
)

type ServiceAuth interface {
	Register(user *UserRegister) error
}

type RepositoryAuth interface {
	Register(user *UserRegister) error
}

type ControllerAuth interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type UserRegister struct {
	Name     string `json:"name" binding:"required" db:"name"`
	Email    string `json:"email" binding:"required" db:"email"`
	Password string `json:"password" binding:"required" db:"password_hash"`
}
