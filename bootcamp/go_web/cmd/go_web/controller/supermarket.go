package controller

import (
	"bootcamp_go_web/internal/supermarket"

	"github.com/gin-gonic/gin"
)

func GetProducts() gin.HandlerFunc {

	return supermarket.GetProducts
}

func GetProductByID() gin.HandlerFunc {
	return supermarket.GetProductByID
}

func SearchProducts() gin.HandlerFunc {
	return supermarket.SearchProducts

}

func AddProduct() gin.HandlerFunc {
	return supermarket.AddProduct
}

func ChangeProduct() gin.HandlerFunc {
	return supermarket.AddProduct
}
