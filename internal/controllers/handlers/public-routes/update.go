package public

import (
	"ice-creams-app/internal/models/domain"
	"ice-creams-app/internal/models/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (hdr *AuthHandler) Update(ctx *gin.Context) {

	req := &dto.UpdateUserRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := &domain.User{
		UserName:      req.UserName,
		Password:      req.Password,
		EmailNew:      req.EmailNew,
		PasswordNew:   req.PasswordNew,
		SecretWordNew: req.SecretWordNew,
	}

	if err := hdr.svc.UpdateUserCredsService(user); err.Error != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User data updated successfully",
	})

}
