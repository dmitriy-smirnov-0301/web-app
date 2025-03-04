package public

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (hdr *AuthHandler) Health(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})

}
