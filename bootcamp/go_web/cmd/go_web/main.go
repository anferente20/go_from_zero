package main

import (
	"bootcamp_go_web/api"
	"bootcamp_go_web/internal/domain/constants"
	"bootcamp_go_web/internal/domain/files"
	"bootcamp_go_web/internal/domain/utils/products"
)

func main() {
	products.GetProducts(files.GetJSONFile(constants.ProductsFile))

	api.SetRoutes()
	api.StartApi()

}
