package prods

import "ice-creams-app/internal/models/domain"

func (svc *IceCreamService) ListIcecreamsService(filter domain.QueryFilter) ([]*domain.IceCream, domain.Error) {

	return svc.repo.ListIcecreams(filter)

}
