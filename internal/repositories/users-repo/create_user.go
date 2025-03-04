package users

import (
	"fmt"
	"ice-creams-app/internal/models/domain"
	"net/http"
)

func (repo *UserRepo) CreateUser(userMdl *domain.User) domain.Error {

	query :=
		`
		INSERT INTO
			users
			(
			user_name,
			email,
			password_hash,
			secret_word_hash
			)
		VALUES
			(
			$1,
			$2,
			$3,
			$4
			)
		RETURNING
			id,
			created_at
		`
	err := repo.db.QueryRow(
		query,
		userMdl.UserName,
		userMdl.Email,
		userMdl.Password,
		userMdl.SecretWord,
	).Scan(
		&userMdl.ID,
		&userMdl.CreatedAt,
	)
	if err != nil {
		log.Warnf("Username or email already exists - %v", err)
		resp.StatusCode = http.StatusConflict
		resp.Error = fmt.Errorf("username or email already exists - %v", err)
		return resp
	}

	log.Infof("User with ID \"%d\" created successfully at %s", userMdl.ID, userMdl.CreatedAt)
	resp.StatusCode = http.StatusCreated
	resp.Error = nil
	return resp

}
