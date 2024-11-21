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
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
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

// Ping
//
//	@Summary		Ping
//	@Description	Ping to check if server is running
//	@Tags			Test v2
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	responses.SuccessResponse
//	@Failure		400	{object}	responses.ErrorResponse
//	@Failure		500	{object}	responses.ErrorResponse
//	@Router			/api/v2/ping [post]
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// ProtectedPing
//
//	@Summary		Protected Ping
//	@Description	Protected Ping to check if token is valid
//	@Tags			Test v2
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Success		200	{object}	responses.SuccessResponse
//	@Failure		400	{object}	responses.ErrorResponse
//	@Failure		500	{object}	responses.ErrorResponse
//	@Router			/api/v2/pingProtected [post]
func ProtectedPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":   "protected pong",
		"userId":    c.GetString("userId"),
		"userEmail": c.GetString("userEmail"),
		"userName":  c.GetString("userName"),
	})
}

func SetupRouter(env *bootstrap.Env, db *sqlx.DB, timeout time.Duration, gin *gin.Engine) {
	gin.Use(CORS())
	router := gin.Group("/api/v2/")

	logger := InitLogger()
	router.Use(middlewares.RequestLoggingMiddleware(logger))

	// Test pings
	router.POST("/ping", Ping)
	router.POST("/pingProtected", middlewares.JwtAuthMiddleware(env.SERVER_SECRET), ProtectedPing)

	if env.ENV_TYPE == "dev" {
		gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// Routers
	NewRouterAuth(env, timeout, db, router)
}
