package public

import (
	"ice-creams-app/internal/models/domain"
	"ice-creams-app/internal/models/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (hdr *AuthHandler) Validate(ctx *gin.Context) {

	req := &dto.ValidateTokenRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token := &domain.Token{
		RefreshToken: req.RefreshToken,
	}

	if err := hdr.svc.ValidateTokenService(token); err.Error != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Error.Error(),
		})
		return
	}

	resp := &dto.ValidateTokenResponse{
		UserName:  token.UserName,
		ExpiresAt: token.ExpiresAt,
		CreatedAt: token.CreatedAt,
		IsValid:   token.IsValid,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": resp,
	})

}
