package middlewares

import (
	"encoding/base64"
	"ice-creams-app/internal/configs"
	"ice-creams-app/internal/pkg/logger"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func BasicAuth(ctx *gin.Context) {

	log := logger.GetLogger()

	config := &configs.AuthConfig{}

	err := config.LoadConfig("AUTH")
	if err != nil {
		log.Fatalf("Failed to load authorization configuration: %v", err)
		ctx.Abort()
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		log.Warn("Missing Authorization header")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	const PREFIX = "Basic "
	if !strings.HasPrefix(authHeader, PREFIX) {
		log.Warn("Invalid Authorization header format")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	encoded := strings.TrimPrefix(authHeader, PREFIX)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		log.Warnf("Failed to decode Authorization header: %v", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	credentials := strings.SplitN(string(decoded), ":", 2)
	if len(credentials) != 2 || credentials[0] != config.Username || credentials[1] != config.Password {
		log.Warn("Invalid credentials")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	log.Info("Authorization successful")
	ctx.Next()

}
