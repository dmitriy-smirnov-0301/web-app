package users

import (
	"fmt"
	"ice-creams-app/internal/models/domain"
	"net/http"
)

func (repo *UserRepo) ReadRefreshToken(tokenMdl *domain.Token) domain.Error {

	query :=
		`
		SELECT
			user_id
		FROM
			refresh_tokens
		WHERE
			token = $1
			AND revoked_at IS NULL
			AND expires_at > CURRENT_TIMESTAMP
		`
	err := repo.db.QueryRow(
		query,
		tokenMdl.RefreshToken,
	).Scan(
		&tokenMdl.UserID,
	)
	if err != nil {
		log.Warnf("Invalid refresh token - %v", err)
		resp.StatusCode = http.StatusUnauthorized
		resp.Error = fmt.Errorf("invalid refresh token - %v", err)
		return resp
	}

	log.Infof("Refresh token read successfully")
	resp.StatusCode = http.StatusOK
	resp.Error = nil
	return resp

}
