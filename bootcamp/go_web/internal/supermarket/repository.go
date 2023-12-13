package supermarket

import (
	"bootcamp_go_web/internal/domain"
	"bootcamp_go_web/internal/supermarket/store"
	"fmt"
)

type ProductRepository interface {
	GetAllProducts() []domain.Product
	GetById(id int) domain.Product
	RegisterProduct(product domain.Product) domain.Product
	UpdateProduct(product domain.Product) domain.Product
	SearchProducts(productPrice float64) []domain.Product
	GetConsumerPrices(ids_list []int) (map[int]domain.Product, float64)
	DeleteProduct(product_id int) error
}

type ProductJSONRepository struct {
	jsonStore store.Store
}

func NewProductJSONRepository(store store.Store) ProductRepository {
	return ProductJSONRepository{store}
}

func (s ProductJSONRepository) GetConsumerPrices(ids_list []int) (map[int]domain.Product, float64) {
	products_list := map[int]domain.Product{}

	if len(ids_list) == 0 {
		ids_list = s.jsonStore.GetAllIds()
	}
	for _, product_id := range ids_list {
		product, is_registered := products_list[product_id]
		prod := s.GetById(product_id)

		if is_registered {

			if product.Quantity+1 <= prod.Quantity {
				product.Quantity++
				products_list[product_id] = product
			}
		} else {
			if prod.Is_published {
				product = prod
				product.Quantity = 1
				products_list[product_id] = product
			}
		}
	}
	var total_price float64

	for _, product := range products_list {
		total_price += product.Price * float64(product.Quantity)
	}

	switch {
	case len(ids_list) < 1 && len(ids_list) > 0:
		total_price *= 1.21
	case len(ids_list) > 10 && len(ids_list) < 20:
		total_price *= 1.17
	case len(ids_list) > 20:
		total_price *= 1.15
	default:
		total_price *= 1.15

	}
	return products_list, total_price

}

func (s ProductJSONRepository) DeleteProduct(product_id int) error {
	return s.jsonStore.Delete(product_id)
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
