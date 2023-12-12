package utils

import (
	"bootcamp_go_web/internal/domain"
	"encoding/json"
	"fmt"
)

func GetProducts(byteValue []byte) []domain.Product {
	products := []domain.Product{}
	json.Unmarshal(byteValue, &products)
	return products
}

func ValidateField(field string, req map[string]any) (any, bool) {
	value, included := req[field]
	isValid := included && fmt.Sprint(value) != ""
	if !isValid {

		return fmt.Sprintf("%s  field is required", field), isValid
	}
	return value, isValid
}

func ValidateCodeVelue(field string, req map[string]any, products []domain.Product) (any, bool) {
	value, isValid := ValidateField(field, req)
	if isValid {
		for _, product := range products {
			if product.Code_value == fmt.Sprint(value) {

				return "code_value field duplicated", false
			}
		}
	}
	return value, isValid
}
