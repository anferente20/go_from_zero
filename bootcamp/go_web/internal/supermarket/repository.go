package supermarket

import (
	"bootcamp_go_web/internal/domain/models"
	"bootcamp_go_web/internal/domain/utils/products"
)

func RegisterProduct(product models.Product) {

	products.Products = append(products.Products, product)
}

func UpdateProduct(product models.Product) {
	for _, prod := range products.Products {
		if product.Id == prod.Id {
			prod.Name = product.Name
			prod.CodeValue = product.CodeValue
			prod.Expiration = product.Expiration
			prod.Price = product.Price
			prod.IsPublished = product.IsPublished
			prod.Quantity = product.Quantity

		}
	}
}
