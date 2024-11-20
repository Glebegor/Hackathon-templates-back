package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

type ginBodyLogger struct {
	gin.ResponseWriter
	body bytes.Buffer
}

func (g *ginBodyLogger) Write(b []byte) (int, error) {
	g.body.Write(b)
	return g.ResponseWriter.Write(b)
}

func RequestLoggingMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ginBodyLogger := &ginBodyLogger{
			body:           bytes.Buffer{},
			ResponseWriter: ctx.Writer,
		}
		ctx.Writer = ginBodyLogger

		var req interface{}
		if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		data, err := json.Marshal(req)
		if err != nil {
			panic(fmt.Errorf("err while marshaling req msg: %v", err))
		}

		ctx.Next()

		logger.WithFields(logrus.Fields{
			"request":  string(data),
			"response": ginBodyLogger.body.String(),
			"status":   ctx.Writer.Status(),
			"method":   ctx.Request.Method,
			"path":     ctx.Request.URL.Path,
			"query":    ctx.Request.URL.RawQuery,
		}).Info("Request logged")
	}
}
