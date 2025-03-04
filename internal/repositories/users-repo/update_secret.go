package users

import (
	"fmt"
	"ice-creams-app/internal/models/domain"
	"net/http"
)

func (repo *UserRepo) UpdateSecret(userMdl *domain.User) domain.Error {

	query :=
		`
		UPDATE
			users
		SET
			password_hash = $1
		WHERE
			user_name = $2
		`
	result, err := repo.db.Exec(
		query,
		userMdl.PasswordNew,
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
		resp.Error = fmt.Errorf("user with name %s not found - %v", userMdl.UserName, err)
		return resp
	}

	log.Infof("User with name \"%s\" updated successfully", userMdl.UserName)
	resp.StatusCode = http.StatusOK
	resp.Error = nil
	return resp

}
