package prods

import (
	"ice-creams-app/internal/models/domain"
	"ice-creams-app/internal/pkg/validator"
	"net/http"
)

func (svc *IceCreamService) UpdateIcecreamService(icecream *domain.IceCream) domain.Error {

	if err = validator.ValidateDates(icecream.ProductionDate, icecream.BestBefore); err != nil {
		log.Error(err)
		resp.StatusCode = http.StatusBadRequest
		resp.Error = err
		return resp
	}

	return svc.repo.UpdateIcecream(icecream)

}
