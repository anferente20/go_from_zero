package products

import (
	"bootcamp_go_web/internal/domain/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Products []models.Product
var ProductID int

func GetProducts(byteValue []byte) {
	json.Unmarshal(byteValue, &Products)
	ProductID = len(Products)
}

func ValidateField(field string, req map[string]any, ctx *gin.Context) (any, bool) {
	value, included := req[field]
	isValid := included && fmt.Sprint(value) != ""
	if !isValid {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("%s  field is required", field),
		})
		return "", isValid
	}
	return value, isValid
}

func ValidateCodeVelue(field string, req map[string]any, ctx *gin.Context) (any, bool) {
	value, isValid := ValidateField(field, req, ctx)
	if isValid {
		for _, product := range Products {
			if product.CodeValue == fmt.Sprint(value) {
				ctx.IndentedJSON(http.StatusBadRequest, gin.H{
					"message": "code_value field duplicated",
				})
				return "", false
			}
		}
	}
	return value, isValid
}
