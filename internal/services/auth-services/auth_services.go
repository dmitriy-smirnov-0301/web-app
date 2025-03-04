package auths

import (
	"ice-creams-app/internal/models/domain"
	"ice-creams-app/internal/pkg/logger"
	"ice-creams-app/internal/repositories/users-repo"
)

type AuthorizationService interface {
	SignupUserService(userMdl *domain.User) domain.Error
	LoginUserService(userMdl *domain.User) (*domain.Token, domain.Error)
	UpdateUserCredsService(userMdl *domain.User) domain.Error
	RecoverPasswordService(userMdl *domain.User) domain.Error
	RefreshTokenService(tokenMdl *domain.Token) domain.Error
	ValidateTokenService(tokenMdl *domain.Token) domain.Error
}

var (
	log  = logger.GetLogger()
	resp = domain.Error{}
	err  error
)

type AuthService struct {
	repo users.UserRepository
}

func NewService(repo users.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}
