package store

import (
	"bootcamp_go_web/internal/domain"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
)

type JsonStore struct {
	FileName string
}

func (js JsonStore) ValidateFile() error {

	jsonFile, err := os.Open(js.FileName)
	defer jsonFile.Close()

	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
func (js JsonStore) Search(id int) (domain.Product, error) {
	jsonFile, err := os.Open(js.FileName)
	defer jsonFile.Close()

	if err != nil {
		return domain.Product{}, errors.New(err.Error())
	}
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return domain.Product{}, errors.New(err.Error())
	}
	products := []domain.Product{}
	product := domain.Product{}
	err = json.Unmarshal(byteValue, &products)
	for _, prod := range products {
		if prod.Id == id {
			product = prod
			break
		}
	}
	return product, nil
}
func (js JsonStore) Modify(product domain.Product) (domain.Product, error) {
	jsonFile, err := os.Open(js.FileName)
	defer jsonFile.Close()

	if err != nil {
		return domain.Product{}, errors.New(err.Error())
	}
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return domain.Product{}, errors.New(err.Error())
	}
	products := []domain.Product{}
	json.Unmarshal(byteValue, &products)
	product_found := false
	updatedProduct := domain.Product{}
	for _, prod := range products {
		if prod.Id == product.Id {
			product_index := slices.Index(products, prod)
			products[product_index] = product
			updatedProduct = products[product_index]
			product_found = true
			break
		}
	}

	file, err := json.MarshalIndent(products, "", " ")
	if err != nil {
		return domain.Product{}, errors.New(err.Error())
	}
	err = os.WriteFile(js.FileName, file, 0644)
	if err != nil {
		return domain.Product{}, errors.New(err.Error())
	}
	if product_found {
		return updatedProduct, nil
	}
	return domain.Product{}, errors.New("Product not found")
}
func (js JsonStore) Delete(id int) error {
	jsonFile, err := os.Open(js.FileName)
	defer jsonFile.Close()

	if err != nil {
		return errors.New(err.Error())
	}
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return errors.New(err.Error())
	}
	products := []domain.Product{}
	json.Unmarshal(byteValue, &products)
	product_found := false
	for _, prod := range products {
		if prod.Id == id {
			product_index := slices.Index(products, prod)
			products = slices.Delete(products, product_index, product_index)
			product_found = true
			break
		}
	}

	file, err := json.MarshalIndent(products, "", " ")
	if err != nil {
		return errors.New(err.Error())
	}
	err = os.WriteFile(js.FileName, file, 0644)
	if err != nil {
		return errors.New(err.Error())
	}
	if product_found {
		return nil
	}
	return errors.New("Product not found")
}

func (js JsonStore) GetAll() []domain.Product {
	jsonFile, err := os.Open(js.FileName)
	defer jsonFile.Close()
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	if err != nil {
		panic(err)
	}
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	products := []domain.Product{}
	err = json.Unmarshal(byteValue, &products)

	if err != nil {
		panic(err)
	}

	return products
}

func (js JsonStore) GetAllIds() []int {
	products := js.GetAll()

	products_ids := []int{}

	for _, prod := range products {
		products_ids = append(products_ids, prod.Id)
	}
	return products_ids
}
func (js JsonStore) Add(product domain.Product) (domain.Product, error) {
	jsonFile, err := os.Open(js.FileName)
	defer jsonFile.Close()

	newProduct := domain.Product{}

	if err != nil {
		return newProduct, err
	}
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return newProduct, err
	}
	products := []domain.Product{}
	err = json.Unmarshal(byteValue, &products)

	if err != nil {
		return newProduct, err
	}

	product.Id = len(products) + 1
	products = append(products, product)
	newProduct = products[len(products)-1]

	file, err := json.MarshalIndent(products, "", " ")
	if err != nil {
		return domain.Product{}, errors.New(err.Error())
	}
	err = os.WriteFile(js.FileName, file, 0644)
	if err != nil {
		return domain.Product{}, errors.New(err.Error())
	}
	return newProduct, nil
}
func (js JsonStore) SearchByPrice(price float64) ([]domain.Product, error) {
	jsonFile, err := os.Open(js.FileName)
	defer jsonFile.Close()
	filtered_products := []domain.Product{}

	if err != nil {
		return filtered_products, errors.New(err.Error())
	}
	fileInfo := json.NewDecoder(jsonFile)
	for {
		var value domain.Product
		if err := fileInfo.Decode(&value); err == io.EOF {
			break // done decoding file
		} else if err != nil {
			return filtered_products, errors.New(err.Error())
		}
		if value.Price >= price {
			filtered_products = append(filtered_products, value)
		}
	}
	return filtered_products, nil
}
