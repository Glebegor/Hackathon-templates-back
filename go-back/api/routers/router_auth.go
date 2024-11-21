package routers

import (
	"project-hackathon/api/controllers"
	"project-hackathon/bootstrap"
	"project-hackathon/repositories"
	"project-hackathon/services"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewRouterAuth(env *bootstrap.Env, timeout time.Duration, db *sqlx.DB, router *gin.RouterGroup) {
	group := router.Group("/auth")

	r := repositories.NewRepositoryAuth(db)
	s := services.NewServiceAuth(r, env, timeout)
	c := controllers.NewControllerAuth(s, env)

	group.POST("/login", c.Login)
	group.POST("/register", c.Register)
	group.POST("/refresh", c.Refresh)

}
