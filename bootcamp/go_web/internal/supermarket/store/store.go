package store

import "bootcamp_go_web/internal/domain"

type Store interface {
	ValidateFile() error
	Search(id int) (domain.Product, error)
	Modify(product domain.Product) (domain.Product, error)
	Delete(id int) error
	GetAll() []domain.Product
	Add(product domain.Product) (domain.Product, error)
	SearchByPrice(price float64) ([]domain.Product, error)
}
