package icecreams

import (
	"fmt"
	"ice-creams-app/internal/models/domain"
	"net/http"
)

func (repo *IceCreamRepo) ListIcecreams(filter domain.QueryFilter) ([]*domain.IceCream, domain.Error) {

	iceCreams := []*domain.IceCream{}

	query := fmt.Sprintf(
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
		ORDER BY
			%s
			%s
		LIMIT
			$1
		OFFSET
			$2
		`,
		filter.SortBy, filter.Order)
	rows, err := repo.db.Query(query, filter.Limit, filter.Offset)
	if err != nil {
		log.Errorf("Failed to fetch ice creams - %v", err)
		resp.StatusCode = http.StatusInternalServerError
		resp.Error = fmt.Errorf("failed to fetch ice creams - %v", err)
		return nil, resp
	}

	defer rows.Close()

	for rows.Next() {
		iceCream := &domain.IceCream{}
		err := rows.Scan(
			&iceCream.ID,
			&iceCream.Name,
			&iceCream.Ingredients,
			&iceCream.ProductionDate,
			&iceCream.BestBefore,
			&iceCream.Price,
			&iceCream.Quantity,
		)
		if err != nil {
			log.Errorf("Failed to scan ice creams - %v", err)
			resp.StatusCode = http.StatusInternalServerError
			resp.Error = fmt.Errorf("failed to scan ice creams - %v", err)
			return nil, resp
		}
		iceCreams = append(iceCreams, iceCream)
	}

	log.Info("Ice creams fetched successfully")
	resp.StatusCode = http.StatusOK
	resp.Error = nil
	return iceCreams, resp

}
