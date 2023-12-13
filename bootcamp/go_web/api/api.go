package api

import (
	"bootcamp_go_web/cmd/go_web/controller"

	"github.com/gin-gonic/gin"
)

var Api = gin.Default()

func SetRoutes() {
	//Create products routes
	productsGroup := Api.Group("/products")
	productRouter := controller.NewProductRouter(productsGroup, true)
	productRouter.ProductRoutes()

}

func StartApi() {
	Api.Run(":8080")

}
