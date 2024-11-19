package routers

import (
	"project-hackathon/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRouter(env *bootstrap.Env, db *sqlx.DB, logger *bootstrap.Logger, timeout time.Duration, gin *gin.Engine) {

}
