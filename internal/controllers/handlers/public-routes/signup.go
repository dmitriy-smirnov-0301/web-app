package public

import (
	"ice-creams-app/internal/models/domain"
	"ice-creams-app/internal/models/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (hdr *AuthHandler) Signup(ctx *gin.Context) {

	req := &dto.SignupUserRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := &domain.User{
		UserName:   req.UserName,
		Email:      req.Email,
		Password:   req.Password,
		SecretWord: req.SecretWord,
	}

	if err := hdr.svc.SignupUserService(user); err.Error != nil {
		ctx.JSON(err.StatusCode, gin.H{
			"error": err.Error.Error(),
		})
		return
	}

	resp := &dto.SignupUserResponse{
		ID:        user.ID,
		UserName:  req.UserName,
		CreatedAt: user.CreatedAt,
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"data":    resp,
	})

}
