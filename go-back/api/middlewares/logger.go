package middlewares

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
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

		rawBody, err := ctx.GetRawData()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  400,
				"message": "Unable to read request body",
			})
			return
		}

		logger.WithFields(logrus.Fields{
			"raw_body": string(rawBody),
		}).Info("Raw body captured")

		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rawBody))

		var req interface{}
		if err := json.Unmarshal(rawBody, &req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  400,
				"message": "Invalid JSON format",
			})
			return
		}

		logger.WithFields(logrus.Fields{
			"request":  string(rawBody),
			"response": ginBodyLogger.body.String(),
			"status":   ctx.Writer.Status(),
			"method":   ctx.Request.Method,
			"path":     ctx.Request.URL.Path,
			"query":    ctx.Request.URL.RawQuery,
		}).Info("Request logged")

		ctx.Next()
	}
}
