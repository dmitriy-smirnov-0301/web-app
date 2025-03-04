package public

import (
	"ice-creams-app/internal/models/domain"
	"ice-creams-app/internal/models/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (hdr *AuthHandler) Login(ctx *gin.Context) {

	req := &dto.LoginUserRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := &domain.User{
		UserName: req.UserName,
		Password: req.Password,
	}

	token, err := hdr.svc.LoginUserService(user)
	if err.Error != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Error.Error(),
		})
		return
	}

	resp := &dto.LoginUserResponse{
		UserName:     req.UserName,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"userName":     resp.UserName,
		"accessToken":  resp.AccessToken,
		"refreshToken": resp.RefreshToken,
	})

}
