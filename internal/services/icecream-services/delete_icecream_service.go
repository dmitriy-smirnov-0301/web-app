package prods

import "ice-creams-app/internal/models/domain"

func (svc *IceCreamService) DeleteIcecreamService(id int) domain.Error {

	return svc.repo.DeleteIcecream(id)

}
