package api

import (
	"bootcamp_go_web/cmd/go_web/controller"
	"bootcamp_go_web/internal/ping"

	"github.com/gin-gonic/gin"
)

var Api = gin.Default()

func SetRoutes() {
	Api.GET("/ping", ping.Ping)
	Api.GET("/products", controller.GetProducts())
	Api.GET("/products/:id", controller.GetProductByID())
	Api.GET("/products/search", controller.SearchProducts())

	Api.POST("/products", controller.AddProduct())

	Api.PUT("/products/:id", controller.ChangeProduct())

}

func StartApi() {
	Api.Run("localhost:8080")

}
