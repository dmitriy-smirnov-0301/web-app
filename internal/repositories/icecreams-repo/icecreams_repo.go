package icecreams

import (
	"database/sql"
	"ice-creams-app/internal/models/domain"
	"ice-creams-app/internal/pkg/logger"
)

type IceCreamRepository interface {
	CreateIcecream(icecream *domain.IceCream) domain.Error
	ListIcecreams(filter domain.QueryFilter) ([]*domain.IceCream, domain.Error)
	ReadIcecream(id int) (*domain.IceCream, domain.Error)
	UpdateIcecream(icecream *domain.IceCream) domain.Error
	DeleteIcecream(id int) domain.Error
}

var (
	log  = logger.GetLogger()
	resp = domain.Error{}
)

type IceCreamRepo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *IceCreamRepo {
	return &IceCreamRepo{
		db: db,
	}
}
