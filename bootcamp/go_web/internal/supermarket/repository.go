package supermarket

import (
	"bootcamp_go_web/internal/domain"
	"bootcamp_go_web/pkg/utils/store"
	"fmt"
)

type ProductRepository struct {
	jsonStore store.Store
}

func NewProductRepository(store store.Store) ProductRepository {
	return ProductRepository{store}
}

func (r *ProductRepository) GetAllProducts() []domain.Product {

	return r.jsonStore.GetAll()
}

func (r *ProductRepository) GetById(id int) domain.Product {
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

func (r *ProductRepository) RegisterProduct(product domain.Product) domain.Product {
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

func (r *ProductRepository) UpdateProduct(product domain.Product) domain.Product {
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

func (r *ProductRepository) SearchProducts(productPrice float64) []domain.Product {
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
