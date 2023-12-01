package gobases3

import (
	"fmt"
)

type Product struct {
	ID          int
	Name        string
	Price       float32
	Description string
	Category    string
}

func (p *Product) Save(product Product, products []Product) []Product {
	return append(products, product)
}

func (p *Product) GetAll(products []Product) {
	for _, product := range products {
		fmt.Println("-------------------------")
		fmt.Println("ID: ", product.ID)
		fmt.Printf("Name: %s \n", product.Name)
		fmt.Println("Description: ", product.Description)
		fmt.Printf("Category: %s \n", product.Category)
	}
}

func GetById(id int, products []Product) Product {
	var product Product
	for _, prod := range products {
		if prod.ID == id {
			product = prod
		}
	}
	return product
}
