package protected

import (
	prods "ice-creams-app/internal/services/icecream-services"
)

type IceCreamHandler struct {
	svc prods.IceCreamsService
}

func NewIceCreamHandler(svc prods.IceCreamsService) *IceCreamHandler {
	return &IceCreamHandler{
		svc: svc,
	}
}
