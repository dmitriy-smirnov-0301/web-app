package auths

import (
	"ice-creams-app/internal/models/domain"
	"ice-creams-app/internal/pkg/hasher"
	"ice-creams-app/internal/pkg/validator"
	"net/http"
)

func (svc *AuthService) SignupUserService(userMdl *domain.User) domain.Error {

	if err = validator.ValidateEmail(userMdl.Email); err != nil {
		log.Warn("Incorrect email format")
		resp.StatusCode = http.StatusBadRequest
		resp.Error = err
		return resp
	}

	if userMdl.Password, err = hasher.Hash(userMdl.Password, 0); err != nil {
		log.Errorf("Failed to hash password - %v", err)
		resp.StatusCode = http.StatusInternalServerError
		resp.Error = err
		return resp
	}

	if userMdl.SecretWord, err = hasher.Hash(userMdl.SecretWord, 0); err != nil {
		log.Errorf("Failed to hash secret word - %v", err)
		resp.StatusCode = http.StatusInternalServerError
		resp.Error = err
		return resp
	}

	return svc.repo.CreateUser(userMdl)

}
