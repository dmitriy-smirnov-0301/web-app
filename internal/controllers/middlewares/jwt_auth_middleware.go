package middlewares

import (
	"ice-creams-app/internal/pkg/enums"
	"ice-creams-app/internal/pkg/jwt"
	"ice-creams-app/internal/pkg/logger"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware(ctx *gin.Context) {

	log := logger.GetLogger()

	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		log.Warn("Missing Authorization header")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	const PREFIX = "Bearer "
	if !strings.HasPrefix(authHeader, PREFIX) {
		log.Warn("Invalid Authorization header format")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	token, err := jwt.ParseToken(strings.TrimPrefix(authHeader, PREFIX))
	if err != nil {
		log.Warnf("Invalid credentials: %v", err)
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}
	if token.TokenType != enums.TokenTypeAccess {
		log.Warn("Invalid token type")
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "Invalid token type",
		})
		return
	}

	log.Info("Authorization successful")
	ctx.Next()

}
