package dbchecker

import (
	"database/sql"
	"errors"
	"ice-creams-app/internal/models/domain"
	"ice-creams-app/internal/pkg/logger"
	"net/http"
)

func CheckDB(db *sql.DB) domain.Error {

	log := logger.GetLogger()
	resp := domain.Error{}

	if db == nil {
		log.Errorf("Database connection status - %v", db)
		resp.StatusCode = http.StatusInternalServerError
		resp.Error = errors.New("database connection is not initialized")
		return resp
	}

	log.Info("Database connection status OK")
	resp.StatusCode = http.StatusOK
	resp.Error = nil
	return resp
}
