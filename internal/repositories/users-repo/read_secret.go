package users

import (
	"fmt"
	"ice-creams-app/internal/models/domain"
	"ice-creams-app/internal/pkg/enums"
	"net/http"
)

func (repo *UserRepo) ReadSecret(userMdl *domain.User, secretType string) domain.Error {

	query := fmt.Sprintf(
		`
		SELECT
			id,
			%s
		FROM
			users
		WHERE
			user_name = $1
		`, secretType)
	switch secretType {
	case enums.SecretTypePassword:
		err := repo.db.QueryRow(
			query,
			userMdl.UserName,
		).Scan(
			&userMdl.ID,
			&userMdl.PasswordHash,
		)
		if err != nil {
			log.Warnf("Invalid username or password - %v", err)
			resp.StatusCode = http.StatusUnauthorized
			resp.Error = fmt.Errorf("invalid username or password - %v", err)
			return resp
		}
	case enums.SecretTypeSecretWord:
		err := repo.db.QueryRow(
			query,
			userMdl.UserName,
		).Scan(
			&userMdl.ID,
			&userMdl.SecretWordHash,
		)
		if err != nil {
			log.Warnf("Invalid username or secret word - %v", err)
			resp.StatusCode = http.StatusUnauthorized
			resp.Error = fmt.Errorf("invalid username or secret word - %v", err)
			return resp
		}
	default:
		log.Error("Incorrect secret type")
		resp.StatusCode = http.StatusInternalServerError
		resp.Error = fmt.Errorf("incorrect secret type")
		return resp
	}

	log.Infof("User secret with name \"%s\" read successfully", userMdl.UserName)
	resp.StatusCode = http.StatusOK
	resp.Error = nil
	return resp

}
