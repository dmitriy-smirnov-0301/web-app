package auths

import (
	"fmt"
	"ice-creams-app/internal/models/domain"
	"ice-creams-app/internal/pkg/enums"
	"ice-creams-app/internal/pkg/hasher"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (svc *AuthService) RecoverPasswordService(userMdl *domain.User) domain.Error {

	if userMdl.PasswordNew, err = hasher.Hash(userMdl.PasswordNew, 0); err != nil {
		log.Errorf("Failed to hash password - %v", err)
		resp.StatusCode = http.StatusInternalServerError
		resp.Error = err
		return resp
	}

	if resp = svc.repo.ReadSecret(userMdl, enums.SecretTypeSecretWord); resp.Error != nil {
		return resp
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userMdl.SecretWordHash), []byte(userMdl.SecretWord)); err != nil {
		log.Warnf("Invalid secret word - %v", err)
		resp.StatusCode = http.StatusUnauthorized
		resp.Error = fmt.Errorf("invalid secret word - %v", err)
		return resp
	}

	return svc.repo.UpdateSecret(userMdl)

}
