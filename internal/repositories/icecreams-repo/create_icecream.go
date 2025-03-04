package icecreams

import (
	"fmt"
	"ice-creams-app/internal/models/domain"
	"net/http"
)

func (repo *IceCreamRepo) CreateIcecream(icecream *domain.IceCream) domain.Error {

	query :=
		`
		INSERT INTO
			icecreams
			(
			name,
			ingredients,
			production_date,
			best_before,
			price,
			quantity
			)
		VALUES
			(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
			)
		RETURNING
			id
		`
	err := repo.db.QueryRow(
		query,
		icecream.Name,
		icecream.Ingredients,
		icecream.ProductionDate,
		icecream.BestBefore,
		icecream.Price,
		icecream.Quantity,
	).Scan(
		&icecream.ID,
	)
	if err != nil {
		log.Errorf("Failed to insert ice cream - %v", err)
		resp.StatusCode = http.StatusInternalServerError
		resp.Error = fmt.Errorf("failed to insert ice cream - %v", err)
		return resp
	}

	log.Infof("Ice cream with ID \"%d\" created successfully", icecream.ID)
	resp.StatusCode = http.StatusOK
	resp.Error = nil
	return resp

}
