package common

import (
	"project-hackathon/core/domain"

	"github.com/gin-gonic/gin"
)

type ServiceAuth interface {
	Register(user *UserRegister) error
	CheckUserByEmailAndPassword(user *UserLogin) (domain.User, error)
	GetUserById(id string) (domain.User, error)
}

type RepositoryAuth interface {
	GetUserById(id string) (domain.User, error)
	Register(user *UserRegister) error
	CheckUserByEmailAndPassword(user *UserLogin) (domain.User, error)
}

type ControllerAuth interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Refresh(ctx *gin.Context)
}

type UserRegister struct {
	Name     string `json:"name" binding:"required" db:"name"`
	Email    string `json:"email" binding:"required" db:"email"`
	Password string `json:"password" binding:"required" db:"password_hash"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required" db:"email"`
	Password string `json:"password" binding:"required" db:"password_hash"`
}
