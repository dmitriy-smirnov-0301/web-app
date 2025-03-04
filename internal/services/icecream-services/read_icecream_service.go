package prods

import "ice-creams-app/internal/models/domain"

func (svc *IceCreamService) ReadIcecreamService(id int) (*domain.IceCream, domain.Error) {

	return svc.repo.ReadIcecream(id)

}
