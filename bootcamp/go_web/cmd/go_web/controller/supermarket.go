package controller

import (
	"bootcamp_go_web/internal/domain"
	"bootcamp_go_web/internal/supermarket"
	"bootcamp_go_web/pkg/constants"
	"bootcamp_go_web/pkg/utils"
	"bootcamp_go_web/pkg/utils/store"
	"fmt"
	"net/http"
	"os"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductRouter struct {
	// grupo al que pertenece este conjunto de urls
	productGroup *gin.RouterGroup
	// el service de productos
	service supermarket.ProductService
}

func NewProductRouter(g *gin.RouterGroup) ProductRouter {
	store := store.JsonStore{FileName: constants.ProductsFile}
	repo := supermarket.NewProductJSONRepository(store)
	service := supermarket.NewProductJSONService(repo)
	return ProductRouter{g, service}
}

func (r *ProductRouter) ProductRoutes() {
	r.productGroup.GET("/", r.GetProducts())
	r.productGroup.GET("/:id", r.GetProductByID())
	r.productGroup.GET("/search", r.SearchProducts())
	r.productGroup.POST("/", r.AddProduct())

	r.productGroup.PATCH("/:id", r.ChangeProduct())
}

func (r *ProductRouter) GetProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data := r.service.GetProducts()
		ctx.JSON(http.StatusOK, data)
	}
}

func (r *ProductRouter) GetProductByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "Bad ID provided",
			})
			return
		}
		product := r.service.GetProductByID(id)
		if product.Id == 0 {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{
				"message": " Product not found",
			})
		} else {
			ctx.IndentedJSON(http.StatusOK, product)
		}
	}
}

func (r *ProductRouter) SearchProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		intPrice, err := strconv.Atoi(ctx.Query("priceGt"))

		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "Bad ID provided",
			})
			return
		}

		price := float64(intPrice)
		response := r.service.SearchProducts(price)

		if len(response) == 0 {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{
				"message": " Products not found",
			})
			return
		}
		ctx.IndentedJSON(http.StatusOK, response)
	}

}

func (r *ProductRouter) AddProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("TOKEN")
		envToken := os.Getenv("TOKEN")

		if token != envToken {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid Token",
			})
			return
		}
		var req map[string]any
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "Bad product provided",
			})
			return
		}

		name, isValid := utils.ValidateField("name", req)
		if !isValid {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": name,
			})
			return
		}
		quantity, isValid := utils.ValidateField("quantity", req)
		if !isValid {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": quantity,
			})
			return
		}
		codeValue, isValid := utils.ValidateCodeVelue("code_value", req, r.service.GetProducts())
		if !isValid {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": codeValue,
			})
			return
		}
		expiration, isValid := utils.ValidateField("expiration", req)
		if !isValid {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": expiration,
			})
			return
		}
		price, isValid := utils.ValidateField("price", req)
		if !isValid {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"price": price,
			})
			return
		}

		qunat, _ := strconv.Atoi(fmt.Sprint(quantity))
		published, _ := strconv.ParseBool(fmt.Sprint(req["is_published"]))
		prix, _ := strconv.ParseFloat(fmt.Sprint(price), 64)
		product := domain.Product{
			Name:         fmt.Sprint(name),
			Quantity:     qunat,
			Code_value:   fmt.Sprint(codeValue),
			Is_published: published,
			Expiration:   fmt.Sprint(expiration),
			Price:        prix,
		}

		insertedProduct := r.service.AddProduct(product)

		ctx.IndentedJSON(http.StatusOK, insertedProduct)
	}
}

func (r *ProductRouter) ChangeProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("TOKEN")
		envToken := os.Getenv("TOKEN")

		if token != envToken {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid Token",
			})
			return
		}
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

		updatedProduct := r.service.GetProductByID(id)
		product_index := slices.Index(r.service.GetProducts(), updatedProduct)

		if product_index == -1 {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{
				"message": " Products not found",
			})
			return
		}
		name, included := req["name"]
		if included {
			updatedProduct.Name = fmt.Sprint(name)
		}
		quantity, included := req["quantity"]
		if included {
			qunat, _ := strconv.Atoi(fmt.Sprint(quantity))
			updatedProduct.Quantity = qunat
		}

		codeValue, included := req["code_value"]
		if included {
			updatedProduct.Code_value = fmt.Sprint(codeValue)
		}

		expiration, included := req["expiration"]
		if included {
			updatedProduct.Expiration = fmt.Sprint(expiration)
		}

		price, included := req["price"]
		if included {
			prix, _ := strconv.ParseFloat(fmt.Sprint(price), 64)
			updatedProduct.Price = prix
		}

		is_published, included := req["is_published"]
		if included {
			published, _ := strconv.ParseBool(fmt.Sprint(is_published))
			updatedProduct.Is_published = published
		}

		updatedProduct = r.service.ChangeProduct(updatedProduct)

		ctx.IndentedJSON(http.StatusOK, updatedProduct)
	}
}
