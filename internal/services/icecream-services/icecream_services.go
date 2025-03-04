package prods

import (
	"ice-creams-app/internal/models/domain"
	"ice-creams-app/internal/pkg/logger"
	"ice-creams-app/internal/repositories/icecreams-repo"
)

type IceCreamsService interface {
	CreateIcecreamService(icecream *domain.IceCream) domain.Error
	ListIcecreamsService(filter domain.QueryFilter) ([]*domain.IceCream, domain.Error)
	ReadIcecreamService(id int) (*domain.IceCream, domain.Error)
	UpdateIcecreamService(icecream *domain.IceCream) domain.Error
	DeleteIcecreamService(id int) domain.Error
}

var (
	log  = logger.GetLogger()
	resp = domain.Error{}
	err  error
)

type IceCreamService struct {
	repo icecreams.IceCreamRepository
}

func NewService(repo icecreams.IceCreamRepository) *IceCreamService {
	return &IceCreamService{
		repo: repo,
	}
}
