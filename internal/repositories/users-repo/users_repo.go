package users

import (
	"database/sql"
	"ice-creams-app/internal/models/domain"
	"ice-creams-app/internal/pkg/logger"
)

type UserRepository interface {
	CreateUser(userMdl *domain.User) domain.Error
	UpdateUser(userMdl *domain.User) domain.Error
	ReadSecret(userMdl *domain.User, secretType string) domain.Error
	UpdateSecret(userMdl *domain.User) domain.Error
	CreateRefreshToken(tokenMdl *domain.Token) domain.Error
	ReadRefreshToken(tokenMdl *domain.Token) domain.Error
	RevokeRefreshToken(tokenStr string) domain.Error
	ValidateRefreshToken(tokenMdl *domain.Token) domain.Error
}

var (
	log  = logger.GetLogger()
	resp = domain.Error{}
)

type UserRepo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}
