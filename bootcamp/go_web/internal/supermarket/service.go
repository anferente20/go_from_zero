package supermarket

import (
	"bootcamp_go_web/internal/domain"
)

type ProductService interface {
	AddProduct(prod domain.Product) domain.Product
	ChangeProduct(prod domain.Product) domain.Product
	SearchProducts(productPrice float64) []domain.Product
	GetProductByID(productId int) domain.Product
	GetProducts() []domain.Product
}
type ProductJSONService struct {
	repository ProductRepository
}

func NewProductJSONService(repo ProductRepository) ProductService {
	return ProductJSONService{repo}
}

func (s ProductJSONService) AddProduct(prod domain.Product) domain.Product {
	return s.repository.RegisterProduct(prod)
}

func (s ProductJSONService) ChangeProduct(prod domain.Product) domain.Product {
	return s.repository.UpdateProduct(prod)
}

func (s ProductJSONService) SearchProducts(productPrice float64) []domain.Product {
	return s.repository.SearchProducts(productPrice)
}

func (s ProductJSONService) GetProductByID(productId int) domain.Product {
	return s.repository.GetById(productId)
}

func (s ProductJSONService) GetProducts() []domain.Product {
	return s.repository.GetAllProducts()

}
