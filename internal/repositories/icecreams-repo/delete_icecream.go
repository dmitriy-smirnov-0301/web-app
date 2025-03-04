package icecreams

import (
	"fmt"
	"ice-creams-app/internal/models/domain"
	"net/http"
)

func (repo *IceCreamRepo) DeleteIcecream(id int) domain.Error {

	query :=
		`
		DELETE FROM
			icecreams
		WHERE
			id = $1
		`
	result, err := repo.db.Exec(query, id)
	if err != nil {
		log.Errorf("Failed to delete ice cream with ID \"%d\" - %v", id, err)
		resp.StatusCode = http.StatusInternalServerError
		resp.Error = fmt.Errorf("failed to delete ice cream with ID '%d' - %v", id, err)
		return resp
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		log.Warnf("Ice cream with ID \"%d\" not found - %v", id, err)
		resp.StatusCode = http.StatusNotFound
		resp.Error = fmt.Errorf("ice cream with ID '%d' not found - %v", id, err)
		return resp
	}

	log.Infof("Ice cream with ID \"%d\" deleted successfully", id)
	resp.StatusCode = http.StatusNoContent
	resp.Error = nil
	return resp

}
