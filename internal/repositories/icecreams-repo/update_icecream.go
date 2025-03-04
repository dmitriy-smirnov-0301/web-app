package icecreams

import (
	"fmt"
	"ice-creams-app/internal/models/domain"
	"net/http"
)

func (repo *IceCreamRepo) UpdateIcecream(icecream *domain.IceCream) domain.Error {

	query :=
		`
		UPDATE
			icecreams 
		SET
			name = $1,
			ingredients = $2,
			production_date = $3,
			best_before = $4,
			price = $5,
			quantity = $6
		WHERE
			id = $7
		`
	result, err := repo.db.Exec(
		query,
		icecream.Name,
		icecream.Ingredients,
		icecream.ProductionDate,
		icecream.BestBefore,
		icecream.Price,
		icecream.Quantity,
		icecream.ID,
	)
	if err != nil {
		log.Errorf("Failed to update ice cream with ID \"%d\" - %v", icecream.ID, err)
		resp.StatusCode = http.StatusInternalServerError
		resp.Error = fmt.Errorf("failed to update ice cream with ID '%d' - %v", icecream.ID, err)
		return resp
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		log.Warnf("Ice cream with ID \"%d\" not found - %v", icecream.ID, err)
		resp.StatusCode = http.StatusNotFound
		resp.Error = fmt.Errorf("ice cream with ID '%d' not found - %v", icecream.ID, err)
		return resp
	}

	log.Infof("Ice cream with ID \"%d\" updated successfully", icecream.ID)
	resp.StatusCode = http.StatusOK
	resp.Error = nil
	return resp

}
