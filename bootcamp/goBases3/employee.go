package gobases3

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	PersonR  Person
	ID       int
	Position string
}

func (e Employee) PrintEmployee() {
	fmt.Println("-------------------------")
	fmt.Printf("Name: %s \n", e.PersonR.Name)
	fmt.Printf("Popsition: %s \n", e.Position)
	fmt.Println("Employee ID: ", e.ID)
	fmt.Println("Person ID: ", e.PersonR.ID)
	fmt.Println("Date of Birth: ", e.PersonR.DateOfBirth)
}
