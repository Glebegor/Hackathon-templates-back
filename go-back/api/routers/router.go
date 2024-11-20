package routers

import (
	"fmt"
	"os"
	"project-hackathon/api/middlewares"
	"project-hackathon/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func InitLogger() *logrus.Logger {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})
	// logger.SetFormatter(&logrus.TextFormatter{
	// 	FullTimestamp: true,
	// })

	logFile, err := os.OpenFile("./logs/logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Errorf("failed to open log file: %v", err))
	}

	logger.SetOutput(logFile)
	logger.SetLevel(logrus.InfoLevel)
	return logger
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func SetupRouter(env *bootstrap.Env, db *sqlx.DB, timeout time.Duration, gin *gin.Engine) {
	gin.Use(CORS())
	router := gin.Group("/api/v2/")

	logger := InitLogger()
	router.Use(middlewares.RequestLoggingMiddleware(logger))

	router.POST("/ping", Ping)

	NewRouterAuth(env, timeout, db, router)
}
