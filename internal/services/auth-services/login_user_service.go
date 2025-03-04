package auths

import (
	"fmt"
	"ice-creams-app/internal/models/domain"
	"ice-creams-app/internal/pkg/enums"
	"ice-creams-app/internal/pkg/jwt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (svc *AuthService) LoginUserService(userMdl *domain.User) (*domain.Token, domain.Error) {

	if resp = svc.repo.ReadSecret(userMdl, enums.SecretTypePassword); resp.Error != nil {
		return nil, resp
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userMdl.PasswordHash), []byte(userMdl.Password)); err != nil {
		log.Warnf("Invalid username or password - %v", err)
		resp.StatusCode = http.StatusUnauthorized
		resp.Error = fmt.Errorf("invalid username or password - %v", err)
		return nil, resp
	}

	accessToken, err := jwt.GenerateToken(userMdl.ID, enums.TokenAccessTTL, enums.TokenTypeAccess)
	if err != nil {
		log.Errorf("Failed to generate access token - %v", err)
		resp.StatusCode = http.StatusInternalServerError
		resp.Error = fmt.Errorf("failed to generate access token - %v", err)
		return nil, resp
	}

	refreshToken, err := jwt.GenerateToken(userMdl.ID, enums.TokenRefreshTTL, enums.TokenTypeRefresh)
	if err != nil {
		log.Errorf("Failed to generate refresh token - %v", err)
		resp.StatusCode = http.StatusInternalServerError
		resp.Error = fmt.Errorf("failed to generate refresh token - %v", err)
		return nil, resp
	}

	token := &domain.Token{
		UserID:       uint32(userMdl.ID),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(enums.TokenRefreshTTL),
		IsValid:      true,
	}

	if resp = svc.repo.CreateRefreshToken(token); resp.Error != nil {
		return nil, resp
	}

	return token, resp

}
