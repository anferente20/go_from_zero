package supermarket

import (
	"bootcamp_go_web/internal/domain"
	"fmt"
	"slices"
)

type ProductRepository struct {
	productsDB []domain.Product
}

func NewProductRepository(prods []domain.Product) ProductRepository {
	return ProductRepository{prods}
}

func (r *ProductRepository) GetAllProducts() []domain.Product {
	fmt.Println(len(r.productsDB))

	return r.productsDB
}

func (r *ProductRepository) GetById(id int) domain.Product {
	product := domain.Product{}
	for _, prod := range r.productsDB {
		if prod.Id == id {
			product = prod
		}
	}
	return product

}

func (r *ProductRepository) RegisterProduct(product domain.Product) domain.Product {
	product.Id = len(r.productsDB) + 1

	r.productsDB = append(r.productsDB, product)

	return product
}

func (r *ProductRepository) UpdateProduct(product domain.Product) domain.Product {
	product_index := 0
	for _, prod := range r.productsDB {
		if product.Id == prod.Id {
			product_index = slices.Index(r.productsDB, prod)
			break
		}
	}
	r.productsDB[product_index] = product
	return r.productsDB[product_index]
}

func (r *ProductRepository) SearchProducts(productPrice float64) []domain.Product {
	productsValid := []domain.Product{}

	for _, product := range r.productsDB {
		if product.Price > productPrice {
			productsValid = append(productsValid, product)
		}
	}
	return productsValid
}
