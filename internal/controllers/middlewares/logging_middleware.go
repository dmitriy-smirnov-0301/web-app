package middlewares

import (
	"ice-creams-app/internal/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger(ctx *gin.Context) {

	start := time.Now()
	ctx.Next()
	duration := time.Since(start)
	status := ctx.Writer.Status()

	log := logger.GetLogger()
	log.WithFields(logrus.Fields{
		"method": ctx.Request.Method,
		"path":   ctx.Request.URL.Path,
		"status": status,
		"time":   duration,
	}).Info("Request completed")

}
