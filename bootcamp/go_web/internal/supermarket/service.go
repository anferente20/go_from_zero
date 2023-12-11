package supermarket

import (
	"bootcamp_go_web/internal/domain/models"
	"bootcamp_go_web/internal/domain/utils/products"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddProduct(ctx *gin.Context) {
	var req map[string]any
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Bad product provided",
		})
		return
	}

	name, isValid := products.ValidateField("name", req, ctx)
	if !isValid {
		return
	}
	quantity, isValid := products.ValidateField("quantity", req, ctx)
	if !isValid {
		return
	}
	codeValue, isValid := products.ValidateCodeVelue("code_value", req, ctx)
	if !isValid {

		return
	}
	expiration, isValid := products.ValidateField("expiration", req, ctx)
	if !isValid {
		return
	}
	price, isValid := products.ValidateField("price", req, ctx)
	if !isValid {
		return
	}
	products.ProductID++

	product := models.Product{
		Id:          products.ProductID,
		Name:        fmt.Sprint(name),
		Quantity:    ctx.GetInt(fmt.Sprint(quantity)),
		CodeValue:   fmt.Sprint(codeValue),
		IsPublished: ctx.GetBool(fmt.Sprint(req["is_published"])),
		Expiration:  fmt.Sprint(expiration),
		Price:       ctx.GetFloat64(fmt.Sprint(price)),
	}

	RegisterProduct(product)

	ctx.IndentedJSON(http.StatusOK, product)
}

func ChangeProduct(ctx *gin.Context) {
	var req map[string]any
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Bad product provided",
		})
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Bad ID provided",
		})
		return
	}

	name, isValid := products.ValidateField("name", req, ctx)
	if !isValid {
		return
	}
	quantity, isValid := products.ValidateField("quantity", req, ctx)
	if !isValid {
		return
	}
	codeValue, isValid := products.ValidateCodeVelue("code_value", req, ctx)
	if !isValid {

		return
	}
	expiration, isValid := products.ValidateField("expiration", req, ctx)
	if !isValid {
		return
	}
	price, isValid := products.ValidateField("price", req, ctx)
	if !isValid {
		return
	}
	products.ProductID++

	product := models.Product{
		Id:          id,
		Name:        fmt.Sprint(name),
		Quantity:    ctx.GetInt(fmt.Sprint(quantity)),
		CodeValue:   fmt.Sprint(codeValue),
		IsPublished: ctx.GetBool(fmt.Sprint(req["is_published"])),
		Expiration:  fmt.Sprint(expiration),
		Price:       ctx.GetFloat64(fmt.Sprint(price)),
	}

	UpdateProduct(product)

	ctx.IndentedJSON(http.StatusOK, product)

}

func SearchProducts(c *gin.Context) {
	intPrice, err := strconv.Atoi(c.Query("priceGt"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Bad ID provided",
		})
		return
	}

	price := float64(intPrice)
	var response []models.Product

	for _, product := range products.Products {
		if product.Price > price {
			response = append(response, product)
		}
	}
	if len(response) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": " Products not found",
		})
	}
	c.IndentedJSON(http.StatusOK, products.Products)
}

func GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Bad ID provided",
		})
		return
	}
	for _, album := range products.Products {
		if album.Id == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": " Product not found",
	})
}

func GetProducts(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, products.Products)
}
