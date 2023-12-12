package supermarket

import (
	"bootcamp_go_web/internal/domain"
)

type ProductService struct {
	repository ProductRepository
}

func NewProductervice(repo ProductRepository) ProductService {
	return ProductService{repo}
}

func (s *ProductService) AddProduct(prod domain.Product) domain.Product {
	return s.repository.RegisterProduct(prod)
}

func (s *ProductService) ChangeProduct(prod domain.Product) domain.Product {
	return s.repository.UpdateProduct(prod)
}

func (s *ProductService) SearchProducts(productPrice float64) []domain.Product {
	return s.repository.SearchProducts(productPrice)
}

func (s *ProductService) GetProductByID(productId int) domain.Product {
	return s.repository.GetById(productId)
}

func (s *ProductService) GetProducts() []domain.Product {
	return s.repository.GetAllProducts()

}
