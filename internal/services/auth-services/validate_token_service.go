package auths

import (
	"fmt"
	"ice-creams-app/internal/models/domain"
	"ice-creams-app/internal/pkg/enums"
	"ice-creams-app/internal/pkg/jwt"
	"net/http"
)

func (svc *AuthService) ValidateTokenService(tokenMdl *domain.Token) domain.Error {

	token, err := jwt.ParseToken(tokenMdl.RefreshToken)
	if err != nil {
		log.Warnf("Invalid token - %v", err)
		resp.StatusCode = http.StatusUnauthorized
		resp.Error = fmt.Errorf("invalid token - %v", err)
		return resp
	}

	tokenMdl.UserID = token.UserID
	tokenMdl.TokenType = token.TokenType
	tokenMdl.ExpiresAt = token.ExpiresAt
	tokenMdl.CreatedAt = token.CreatedAt
	tokenMdl.IsValid = token.IsValid

	if tokenMdl.TokenType == enums.TokenTypeAccess {
		log.Warnf("Invalid token type")
		resp.StatusCode = http.StatusBadRequest
		resp.Error = fmt.Errorf("invalid token type")
		return resp
	}

	return svc.repo.ValidateRefreshToken(tokenMdl)

}
