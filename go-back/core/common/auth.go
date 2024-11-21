package common

import (
	"project-hackathon/core/domain"

	"github.com/gin-gonic/gin"
)

type ServiceAuth interface {
	Register(user *UserRegister) error
	Login(user *UserLogin) (string, error)
}

type RepositoryAuth interface {
	Register(user *UserRegister) error
	CheckUserByEmailAndPassword(user *UserLogin) (domain.User, error)
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

type UserLogin struct {
	Email    string `json:"email" binding:"required" db:"email"`
	Password string `json:"password" binding:"required" db:"password_hash"`
}

type JWTClaims struct {
	Email string `json:"email" db:"email"`
	Name  string `json:"name" db:"name"`
	Id    int    `json:"id" db:"id"`
	Exp   int64  `json:"exp"`
}
