package auths

import (
	"fmt"
	"ice-creams-app/internal/models/domain"
	"ice-creams-app/internal/pkg/enums"
	"ice-creams-app/internal/pkg/jwt"
	"net/http"
	"time"
)

func (svc *AuthService) RefreshTokenService(tokenMdl *domain.Token) domain.Error {

	if resp = svc.repo.ReadRefreshToken(tokenMdl); resp.Error != nil {
		return resp
	}

	if resp = svc.repo.RevokeRefreshToken(tokenMdl.RefreshToken); resp.Error != nil {
		return resp
	}

	newAccessToken, err := jwt.GenerateToken(tokenMdl.UserID, enums.TokenAccessTTL, enums.TokenTypeAccess)
	if err != nil {
		log.Errorf("Failed to generate access token - %v", err)
		resp.StatusCode = http.StatusInternalServerError
		resp.Error = fmt.Errorf("failed to generate access token - %v", err)
		return resp
	}

	newRefreshToken, err := jwt.GenerateToken(tokenMdl.UserID, enums.TokenRefreshTTL, enums.TokenTypeRefresh)
	if err != nil {
		log.Errorf("Failed to generate refresh token - %v", err)
		resp.StatusCode = http.StatusInternalServerError
		resp.Error = fmt.Errorf("failed to generate refresh token - %v", err)
		return resp
	}

	tokenMdl.AccessToken = newAccessToken
	tokenMdl.RefreshToken = newRefreshToken
	tokenMdl.ExpiresAt = time.Now().Add(enums.TokenRefreshTTL)

	return svc.repo.CreateRefreshToken(tokenMdl)

}
