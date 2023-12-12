package supermarket

import (
	"bootcamp_go_web/internal/domain"
	"bootcamp_go_web/pkg/utils/store"
	"fmt"
)

type ProductRepository interface {
	GetAllProducts() []domain.Product
	GetById(id int) domain.Product
	RegisterProduct(product domain.Product) domain.Product
	UpdateProduct(product domain.Product) domain.Product
	SearchProducts(productPrice float64) []domain.Product
}

type ProductJSONRepository struct {
	jsonStore store.Store
}

func NewProductJSONRepository(store store.Store) ProductRepository {
	return ProductJSONRepository{store}
}

func (r ProductJSONRepository) GetAllProducts() []domain.Product {

	return r.jsonStore.GetAll()
}

func (r ProductJSONRepository) GetById(id int) domain.Product {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	product, err := r.jsonStore.Search(id)
	if err != nil {
		panic(err)
	}
	return product

}

func (r ProductJSONRepository) RegisterProduct(product domain.Product) domain.Product {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	prod, err := r.jsonStore.Add(product)
	if err != nil {
		panic(err)
	}
	return prod
}

func (r ProductJSONRepository) UpdateProduct(product domain.Product) domain.Product {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	prod, err := r.jsonStore.Modify(product)
	if err != nil {
		panic(err)
	}
	return prod

}

func (r ProductJSONRepository) SearchProducts(productPrice float64) []domain.Product {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	productsValid, err := r.jsonStore.SearchByPrice(productPrice)
	if err != nil {
		panic(err)
	}

	return productsValid
}
