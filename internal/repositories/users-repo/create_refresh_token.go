package users

import (
	"fmt"
	"ice-creams-app/internal/models/domain"
	"net/http"
)

func (repo *UserRepo) CreateRefreshToken(tokenMdl *domain.Token) domain.Error {

	query :=
		`
		INSERT INTO
			refresh_tokens
			(
			user_id,
			token,
			expires_at
			)
		VALUES
			(
			$1,
			$2,
			$3
			)
		`
	_, err := repo.db.Exec(
		query,
		tokenMdl.UserID,
		tokenMdl.RefreshToken,
		tokenMdl.ExpiresAt,
	)
	if err != nil {
		log.Errorf("Failed to save refresh token - %v", err)
		resp.StatusCode = http.StatusInternalServerError
		resp.Error = fmt.Errorf("failed to save refresh token - %v", err)
		return resp
	}

	log.Infof("Tokens added successfully:\naccess - %s\nrefresh - %s", tokenMdl.AccessToken, tokenMdl.RefreshToken)
	resp.StatusCode = http.StatusOK
	resp.Error = nil
	return resp

}
