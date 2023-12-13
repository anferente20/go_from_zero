package tests

import (
	"bootcamp_go_web/cmd/go_web/controller"
	"bootcamp_go_web/pkg/utils/web"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "1q2w3e4r*")
	var api = gin.Default()
	productsGroup := api.Group("/products")
	productRouter := controller.NewProductRouter(productsGroup, false)
	productRouter.ProductRoutes()
	return api
}

func createRequestTest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", "1q2w3e4r*")
	return req, httptest.NewRecorder()
}

func TestGetProducts(t *testing.T) {
	server := createServer()
	t.Run("Validate get all productss SUCCESS",
		func(t *testing.T) {
			req, rr := createRequestTest(http.MethodGet, "/products/", "")
			server.ServeHTTP(rr, req)

			assert.Equal(t, 200, rr.Code)

			var objRes web.Response

			err := json.Unmarshal(rr.Body.Bytes(), &objRes)
			products := objRes.Data.([]interface{})
			assert.Nil(t, err)
			assert.True(t, len(products) > 0)
		})
}

func TestAddProduct(t *testing.T) {
	server := createServer()
	t.Run("Validate add product SUCCESS",
		func(t *testing.T) {
			req, rr := createRequestTest(http.MethodPost, "/products/", `{"name":"Tester","quantity":1,"code_value":"S7304A","is_published":true,"expiration":"02/04/2021","price":976.27}`)

			server.ServeHTTP(rr, req)

			assert.Equal(t, 200, rr.Code)

			var objRes web.Response

			err := json.Unmarshal(rr.Body.Bytes(), &objRes)
			assert.Nil(t, err)

			assert.True(t, objRes.Data != nil)
		})
}
