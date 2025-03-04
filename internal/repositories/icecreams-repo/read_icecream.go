package icecreams

import (
	"database/sql"
	"fmt"
	"ice-creams-app/internal/models/domain"
	"net/http"
)

func (repo *IceCreamRepo) ReadIcecream(id int) (*domain.IceCream, domain.Error) {

	iceCream := &domain.IceCream{}

	query :=
		`
		SELECT
			id,
			name,
			ingredients,
			production_date,
			best_before,
			price,
			quantity
		FROM
			icecreams 
		WHERE
			id = $1
		`
	err := repo.db.QueryRow(
		query,
		id,
	).Scan(
		&iceCream.ID,
		&iceCream.Name,
		&iceCream.Ingredients,
		&iceCream.ProductionDate,
		&iceCream.BestBefore,
		&iceCream.Price,
		&iceCream.Quantity,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Warnf("Ice cream with ID \"%d\" not found - %v", id, err)
			resp.StatusCode = http.StatusNotFound
			resp.Error = fmt.Errorf("ice cream with ID '%d' not found - %v", id, err)
			return nil, resp
		}
		log.Errorf("Failed to fetch ice cream with ID \"%d\" - %v", id, err)
		resp.StatusCode = http.StatusInternalServerError
		resp.Error = fmt.Errorf("failed to fetch ice cream with ID '%d' - %v", id, err)
		return nil, resp
	}

	log.Infof("Ice cream with ID \"%d\" fetched successfully", id)
	resp.StatusCode = http.StatusOK
	resp.Error = nil
	return iceCream, resp

}
