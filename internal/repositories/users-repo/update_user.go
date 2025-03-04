package users

import (
	"fmt"
	"ice-creams-app/internal/models/domain"
	"net/http"
)

func (repo *UserRepo) UpdateUser(userMdl *domain.User) domain.Error {

	query :=
		`
		UPDATE
			users
		SET
			email = $1,
			password_hash = $2,
			secret_word_hash = $3
		WHERE
			user_name = $4
		`
	result, err := repo.db.Exec(
		query,
		userMdl.EmailNew,
		userMdl.PasswordNew,
		userMdl.SecretWordNew,
		userMdl.UserName,
	)
	if err != nil {
		log.Errorf("Failed to update user with name \"%s\" - %v", userMdl.UserName, err)
		resp.StatusCode = http.StatusInternalServerError
		resp.Error = fmt.Errorf("failed to update user with name %s - %v", userMdl.UserName, err)
		return resp
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		log.Warnf("User with name \"%s\" not found - %v", userMdl.UserName, err)
		resp.StatusCode = http.StatusNotFound
		resp.Error = fmt.Errorf("user not found with name %s - %v", userMdl.UserName, err)
		return resp
	}

	log.Infof("User with name \"%s\" updated successfully", userMdl.UserName)
	resp.StatusCode = http.StatusOK
	resp.Error = nil
	return resp

}
