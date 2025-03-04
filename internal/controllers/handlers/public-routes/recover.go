package public

import (
	"ice-creams-app/internal/models/domain"
	"ice-creams-app/internal/models/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (hdr *AuthHandler) Recover(ctx *gin.Context) {

	req := &dto.RecoverUserRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := &domain.User{
		UserName:    req.UserName,
		SecretWord:  req.SecretWord,
		PasswordNew: req.PasswordNew,
	}

	if err := hdr.svc.RecoverPasswordService(user); err.Error != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User data recovered successfully",
	})

}
