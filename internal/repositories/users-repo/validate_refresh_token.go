package users

import (
	"fmt"
	"ice-creams-app/internal/models/domain"
	"net/http"
)

func (repo *UserRepo) ValidateRefreshToken(tokenMdl *domain.Token) domain.Error {

	query :=
		`
		SELECT
			user_name
		FROM
			users
		WHERE
			id = $1
		`
	err := repo.db.QueryRow(
		query,
		tokenMdl.UserID,
	).Scan(
		&tokenMdl.UserName,
	)
	if err != nil {
		log.Warnf("User not found - %v", err)
		resp.StatusCode = http.StatusNotFound
		resp.Error = fmt.Errorf("user not found - %v", err)
		return resp
	}

	log.Infof("Token \"%s\" is valid", tokenMdl.RefreshToken)
	resp.StatusCode = http.StatusOK
	resp.Error = nil
	return resp

}
