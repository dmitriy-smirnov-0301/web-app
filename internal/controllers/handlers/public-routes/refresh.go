package public

import (
	"ice-creams-app/internal/models/domain"
	"ice-creams-app/internal/models/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (hdr *AuthHandler) Refresh(ctx *gin.Context) {

	req := &dto.RefreshTokenRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token := &domain.Token{
		RefreshToken: req.RefreshToken,
	}

	if err := hdr.svc.RefreshTokenService(token); err.Error != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Error.Error(),
		})
		return
	}

	resp := &dto.RefreshTokenResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"accessToken":  resp.AccessToken,
		"refreshToken": resp.RefreshToken,
	})

}
