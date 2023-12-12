package controller

import (
	"bootcamp_go_web/internal/domain"
	"bootcamp_go_web/internal/supermarket"
	"bootcamp_go_web/pkg/constants"
	"bootcamp_go_web/pkg/utils"
	"bootcamp_go_web/pkg/utils/store"
	"bootcamp_go_web/pkg/utils/web"
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

		resp := web.Response{Data: data}
		web.SetResponse(resp, ctx)
	}
}

func (r *ProductRouter) GetProductByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			resp := web.ErrorResponse{
				Status:  http.StatusBadRequest,
				Code:    constants.BadRequestCode,
				Message: constants.BadRequestMessage,
			}
			web.SetErrorResponse(resp, ctx)

			return
		}
		product := r.service.GetProductByID(id)
		if product.Id == 0 {
			resp := web.ErrorResponse{
				Status:  http.StatusNotFound,
				Code:    constants.NotFoundCode,
				Message: constants.NotFoundMessage,
			}
			web.SetErrorResponse(resp, ctx)

		} else {
			resp := web.Response{Data: product}
			web.SetResponse(resp, ctx)
		}
	}
}

func (r *ProductRouter) SearchProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		intPrice, err := strconv.Atoi(ctx.Query("priceGt"))

		if err != nil {
			resp := web.ErrorResponse{
				Status:  http.StatusBadRequest,
				Code:    constants.BadRequestCode,
				Message: constants.BadRequestMessage,
			}
			web.SetErrorResponse(resp, ctx)
			return
		}

		price := float64(intPrice)
		response := r.service.SearchProducts(price)

		if len(response) == 0 {
			resp := web.ErrorResponse{
				Status:  http.StatusNotFound,
				Code:    constants.NotFoundCode,
				Message: constants.NotFoundMessage,
			}
			web.SetErrorResponse(resp, ctx)

			return
		}
		resp := web.Response{Data: response}
		web.SetResponse(resp, ctx)
	}

}

func (r *ProductRouter) AddProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("TOKEN")
		envToken := os.Getenv("TOKEN")

		if token != envToken {
			resp := web.ErrorResponse{
				Status:  http.StatusUnauthorized,
				Code:    constants.UnauthorizedCode,
				Message: constants.UnauthorizedMessage,
			}
			web.SetErrorResponse(resp, ctx)

			return
		}
		var req map[string]any
		if err := ctx.ShouldBind(&req); err != nil {
			resp := web.ErrorResponse{
				Status:  http.StatusBadRequest,
				Code:    constants.BadRequestCode,
				Message: constants.BadRequestMessage,
			}
			web.SetErrorResponse(resp, ctx)
			return
		}

		name, isValid := utils.ValidateField("name", req)
		if !isValid {
			resp := web.ErrorResponse{
				Status:  http.StatusBadRequest,
				Code:    constants.BadRequestCode,
				Message: fmt.Sprint(name),
			}
			web.SetErrorResponse(resp, ctx)

			return
		}
		quantity, isValid := utils.ValidateField("quantity", req)
		if !isValid {
			resp := web.ErrorResponse{
				Status:  http.StatusBadRequest,
				Code:    constants.BadRequestCode,
				Message: fmt.Sprint(quantity),
			}
			web.SetErrorResponse(resp, ctx)
			return
		}
		codeValue, isValid := utils.ValidateCodeVelue("code_value", req, r.service.GetProducts())
		if !isValid {
			resp := web.ErrorResponse{
				Status:  http.StatusBadRequest,
				Code:    constants.BadRequestCode,
				Message: fmt.Sprint(codeValue),
			}
			web.SetErrorResponse(resp, ctx)
			return
		}
		expiration, isValid := utils.ValidateField("expiration", req)
		if !isValid {
			resp := web.ErrorResponse{
				Status:  http.StatusBadRequest,
				Code:    constants.BadRequestCode,
				Message: fmt.Sprint(expiration),
			}
			web.SetErrorResponse(resp, ctx)
			return
		}
		price, isValid := utils.ValidateField("price", req)
		if !isValid {
			resp := web.ErrorResponse{
				Status:  http.StatusBadRequest,
				Code:    constants.BadRequestCode,
				Message: fmt.Sprint(price),
			}
			web.SetErrorResponse(resp, ctx)
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
		resp := web.Response{Data: insertedProduct}
		web.SetResponse(resp, ctx)
	}
}

func (r *ProductRouter) ChangeProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("TOKEN")
		envToken := os.Getenv("TOKEN")

		if token != envToken {
			resp := web.ErrorResponse{
				Status:  http.StatusUnauthorized,
				Code:    constants.UnauthorizedCode,
				Message: constants.UnauthorizedMessage,
			}
			web.SetErrorResponse(resp, ctx)
			return
		}
		var req map[string]any
		if err := ctx.ShouldBind(&req); err != nil {
			resp := web.ErrorResponse{
				Status:  http.StatusBadRequest,
				Code:    constants.BadRequestCode,
				Message: constants.BadRequestMessage,
			}
			web.SetErrorResponse(resp, ctx)
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			resp := web.ErrorResponse{
				Status:  http.StatusBadRequest,
				Code:    constants.BadRequestCode,
				Message: constants.BadRequestMessage,
			}
			web.SetErrorResponse(resp, ctx)
			return
		}

		updatedProduct := r.service.GetProductByID(id)
		product_index := slices.Index(r.service.GetProducts(), updatedProduct)

		if product_index == -1 {
			resp := web.ErrorResponse{
				Status:  http.StatusNotFound,
				Code:    constants.NotFoundCode,
				Message: constants.NotFoundMessage,
			}
			web.SetErrorResponse(resp, ctx)
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
		resp := web.Response{Data: updatedProduct}
		web.SetResponse(resp, ctx)
		ctx.IndentedJSON(http.StatusOK, updatedProduct)
	}
}
