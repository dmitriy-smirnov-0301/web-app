package auths

import (
	"fmt"
	"ice-creams-app/internal/models/domain"
	"ice-creams-app/internal/pkg/enums"
	"ice-creams-app/internal/pkg/hasher"
	"ice-creams-app/internal/pkg/validator"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (svc *AuthService) UpdateUserCredsService(userMdl *domain.User) domain.Error {

	if err = validator.ValidateEmail(userMdl.EmailNew); err != nil {
		log.Warn("Incorrect email format")
		resp.StatusCode = http.StatusBadRequest
		resp.Error = err
		return resp
	}

	if userMdl.PasswordNew, err = hasher.Hash(userMdl.PasswordNew, 0); err != nil {
		log.Errorf("Failed to hash password - %v", err)
		resp.StatusCode = http.StatusInternalServerError
		resp.Error = err
		return resp
	}

	if userMdl.SecretWordNew, err = hasher.Hash(userMdl.SecretWordNew, 0); err != nil {
		log.Errorf("Failed to hash secret word - %v", err)
		resp.StatusCode = http.StatusInternalServerError
		resp.Error = err
		return resp
	}

	if resp = svc.repo.ReadSecret(userMdl, enums.SecretTypePassword); resp.Error != nil {
		return resp
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userMdl.PasswordHash), []byte(userMdl.Password)); err != nil {
		log.Warnf("Invalid password - %v", err)
		resp.StatusCode = http.StatusUnauthorized
		resp.Error = fmt.Errorf("invalid password - %v", err)
		return resp
	}

	return svc.repo.UpdateUser(userMdl)

}
