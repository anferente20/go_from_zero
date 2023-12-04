package gobases4

import (
	"errors"
	"fmt"
	"os"
)

type InvalidFieldError struct {
	Err  error
	Code int
}

func (m *InvalidFieldError) Error() string {
	return fmt.Sprintf("%s, Status code: %d", m.Err, m.Code)
}

type Customer struct {
	File  string
	Name  string
	ID    int
	Phone string
	Home  string
}

func (c Customer) toString() string {
	return fmt.Sprintf("%d - %s - %s - %s - %s", c.ID, c.File, c.Name, c.Home, c.Phone)
}

func Read(fileName string) string {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	data, err := os.ReadFile(fileName)

	fmt.Println("Reading Finished")
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}
	return string(data)
}

func Write(filename string, newCustomer Customer, customers []Customer) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	validateUserExistence(newCustomer, customers)
	_, err := validateCustomerFields(newCustomer)
	if err != nil {
		panic("Several errors were detected at runtime 1")
	}

	file, err2 := os.Create(filename)
	defer file.Close()

	if err2 != nil {
		panic("Several errors were detected at runtime 2")
	}
	_, err3 := file.WriteString(newCustomer.toString())
	if err3 != nil {
		panic(err3.Error())
	}

}

func validateUserExistence(newCustomer Customer, customers []Customer) bool {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	customerExists := false
	for _, customer := range customers {
		if customer.ID == newCustomer.ID {
			customerExists = true
			break
		}
	}
	if customerExists {
		panic("Error: client already exists")
	}
	return customerExists
}

func validateCustomerFields(newCustomer Customer) (bool, error) {
	if newCustomer.File == "" {
		return false, &InvalidFieldError{Code: 1, Err: errors.New("File Field is not valid")}
	}
	if newCustomer.Name == "" {
		return false, &InvalidFieldError{Code: 2, Err: errors.New("Name Field is not valid")}
	}
	if newCustomer.Phone == "" {
		return false, &InvalidFieldError{Code: 3, Err: errors.New("Phone Field is not valid")}
	}
	if newCustomer.Home == "" {
		return false, &InvalidFieldError{Code: 4, Err: errors.New("Home Field is not valid")}
	}
	if newCustomer.ID == 0 {
		return false, &InvalidFieldError{Code: 4, Err: errors.New("ID Field is not valid")}
	}

	return true, nil
}
