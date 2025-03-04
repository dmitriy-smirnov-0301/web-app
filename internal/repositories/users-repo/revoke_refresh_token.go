package users

import (
	"fmt"
	"ice-creams-app/internal/models/domain"
	"net/http"
)

func (repo *UserRepo) RevokeRefreshToken(tokenStr string) domain.Error {

	query :=
		`
		UPDATE
			refresh_tokens
		SET
			revoked_at = CURRENT_TIMESTAMP
		WHERE
			token = $1
		`
	_, err := repo.db.Exec(
		query,
		tokenStr,
	)
	if err != nil {
		log.Errorf("Failed to update refresh token - %v", err)
		resp.StatusCode = http.StatusInternalServerError
		resp.Error = fmt.Errorf("failed to update refresh token - %v", err)
		return resp
	}

	log.Infof("Refresh token revoked successfully")
	resp.StatusCode = http.StatusOK
	resp.Error = nil
	return resp

}
