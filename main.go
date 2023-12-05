package main

import (
	"fmt"
	"time"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Println("Hello world")

	// variables.ShowIntegers()
	// variables.OtherVariables()

	// fmt.Println("--------------------------------")
	// fmt.Println("----------- Cast Integer -----------")
	// var state, text = variables.CastToText(190919)
	// fmt.Println("State = ", state)
	// fmt.Println("Text = ", text)
	// fmt.Println("--------------------------------")

	// fmt.Println("--------------------------------")
	// fmt.Println("----------- EX 1 -----------")
	// fmt.Println(exercices.IsMoreThan100("a "))
	// fmt.Println("--------------------------------")

	// fmt.Println("--------------------------------")
	// fmt.Println("----------- Inputs -----------")
	// inputs.GetNumbers()
	// fmt.Println("--------------------------------")

	// fmt.Println("--------------------------------")
	// fmt.Println("----------- EX 2 -----------")
	// exercices.GetMultiplicationTable()
	// fmt.Println("--------------------------------")

	// fmt.Println("--------------------------------")
	// fmt.Println("-----------Anonymous functions -----------")
	// functions.Calcules(10, 15)
	// fmt.Println("--------------------------------")
	// fmt.Println("-----------Closures -----------")
	// functions.CallClosure(2)
	// fmt.Println("--------------------------------")
	// fmt.Println("-----------Recursion -----------")
	// fmt.Println("The summatory of 10 is: ", functions.Summatory(10))
	// fmt.Println("--------------------------------")

	// fmt.Println("--------------------------------")
	// fmt.Println("----------- Letters -----------")
	// gobases1.GetLetters()
	// fmt.Println("--------------------------------")
	// fmt.Println("----------- Loans -----------")
	// gobases1.CanBorrow()
	// fmt.Println("--------------------------------")
	// fmt.Println("--------------------------------")
	// fmt.Println("----------- Salary -----------")
	// gobases1.GetSalary()
	// fmt.Println("--------------------------------")

	// Grades
	// var grades []int
	// var err error
	// var repeat int
	// var grade int
	// scanner := bufio.NewScanner(os.Stdin)
	// for {
	// 	fmt.Println("enter ", len(grades), " grade: ")
	// 	if scanner.Scan() {
	// 		grade, err = strconv.Atoi(scanner.Text())
	// 		if err != nil {
	// 			fmt.Println("Wrong number.")
	// 		}
	// 		grades = append(grades, grade)
	// 	}
	// 	fmt.Println("do you wanna get another grade? (1 - yes): ")
	// 	if scanner.Scan() {
	// 		repeat, err = strconv.Atoi(scanner.Text())
	// 		if err != nil || repeat != 1 {
	// 			fmt.Println("Thanks.")
	// 			break
	// 		}
	// 	}
	// }
	// gobases1.CalculateGrade(grades...)

	//Salary day 2
	// const (
	// 	categoryA = "A"
	// 	categoryB = "B"
	// 	categoryC = "C"
	// )

	// var err error
	// var hours int
	// var category string
	// var repeat int
	// fmt.Println("--------------------------------")
	// for {
	// 	fmt.Println("Enter your Employee category: ")
	// 	if scanner.Scan() {
	// 		category = scanner.Text()
	// 		if category != categoryA && category != categoryB && category != categoryC {
	// 			fmt.Println("Wrong category.")
	// 		} else {
	// 			for {
	// 				fmt.Println("Enter the hours your employee worked this month: ")
	// 				if scanner.Scan() {
	// 					hours, err = strconv.Atoi(scanner.Text())
	// 					if err != nil {
	// 						fmt.Println("Wrong Number.")
	// 					} else {
	// 						fmt.Println("Employee salary: ", gobases2.GetSalary(hours, category))
	// 						break
	// 					}

	// 				}

	// 			}
	// 		}

	// 		fmt.Println("do you wanna get another salary? (1 - yes): ")
	// 		if scanner.Scan() {
	// 			repeat, err = strconv.Atoi(scanner.Text())
	// 			if err != nil || repeat != 1 {
	// 				fmt.Println("Thanks.")
	// 				break

	// 			}
	// 		}
	// 	}

	// }
	// fmt.Println("--------------------------------")

	// Stats
	// fmt.Println("--------------------------------")
	// var grades []float32
	// var err error
	// var repeat int
	// var grade int
	// for {
	// 	fmt.Println("enter ", len(grades), " grade: ")
	// 	if scanner.Scan() {
	// 		grade, err = strconv.Atoi(scanner.Text())
	// 		if err != nil {
	// 			fmt.Println("Wrong number.")
	// 		}
	// 		grades = append(grades, float32(grade))
	// 	}
	// 	fmt.Println("do you wanna get another grade? (1 - yes): ")
	// 	if scanner.Scan() {
	// 		repeat, err = strconv.Atoi(scanner.Text())
	// 		if err != nil || repeat != 1 {
	// 			fmt.Println("Thanks.")
	// 			break
	// 		}
	// 	}
	// }
	// minimun := gobases2.GetFunction("minimum")
	// minValue, err := minimun(grades...)
	// fmt.Println("The minimum is: ", minValue, err)

	// maximum := gobases2.GetFunction("maximum")
	// maxValue, err := maximum(grades...)
	// fmt.Println("The maximum is: ", maxValue, err)

	// average := gobases2.GetFunction("average")
	// aveValue, err := average(grades...)
	// fmt.Println("The average is: ", aveValue, err)

	// fmt.Println("--------------------------------")
	// fmt.Println("--------------------------------")
	// Products

	// var products []gobases3.Product
	// product1 := gobases3.Product{
	// 	ID:          1,
	// 	Name:        "Carnecita",
	// 	Description: "Carne fresca de velociraptor",
	// 	Category:    "Comida",
	// }
	// products = product1.Save(product1, products)

	// product1.GetAll(products)

	// Employee
	// person := gobases3.Person{
	// 	ID:          1,
	// 	Name:        "Andres F",
	// 	DateOfBirth: "05/04/1997",
	// }

	// employee := gobases3.Employee{
	// 	ID:       1,
	// 	Position: "Software Developer",
	// 	PersonR:  person,
	// }
	// employee.PrintEmployee()

	//Peoduc type
	// smallProduct := gobases3.NewProduct(gobases3.Small, 2000)
	// mediumlProduct := gobases3.NewProduct(gobases3.Medium, 5000)
	// largeProduct := gobases3.NewProduct(gobases3.Large, 20000)

	// fmt.Println("Small product: ", smallProduct.GetPrice())
	// fmt.Println("Medium product: ", mediumlProduct.GetPrice())
	// fmt.Println("Large product: ", largeProduct.GetPrice())

	//Salary
	// salary, err := gobases4.CalculateSalary(30, 2000)
	// var errLow *gobases4.LowSalaryError
	// var errTax *gobases4.TaxableError

	// if err != nil && errors.As(err, &errLow) {
	// 	fmt.Printf("Salary: %g, LowSalaryError: %s \n", salary, err.Error())
	// } else if err != nil && errors.As(err, &errTax) {
	// 	fmt.Printf("Salary: %g, TaxableError: %s \n", salary, err.Error())
	// } else {
	// 	fmt.Printf("Salary: %g \n", salary)
	// }
	// fmt.Println("--------------------------------")

	// salary, err = gobases4.CalculateSalary(80, 100)
	// if err != nil && errors.As(err, &errLow) {
	// 	fmt.Printf("Salary: %g, LowSalaryError: %s \n", salary, err.Error())
	// } else if err != nil && errors.As(err, &errTax) {
	// 	fmt.Printf("Salary: %g, TaxableError: %s \n", salary, err.Error())
	// } else {
	// 	fmt.Printf("Salary: %g \n", salary)
	// }
	// fmt.Println("--------------------------------")

	// salary, err = gobases4.CalculateSalary(120, 1000)
	// if err != nil && errors.As(err, &errLow) {
	// 	fmt.Printf("Salary: %g, LowSalaryError: %s \n", salary, err.Error())
	// } else if err != nil && errors.As(err, &errTax) {
	// 	fmt.Printf("Salary: %g, TaxableError: %s \n", salary, err.Error())
	// } else {
	// 	fmt.Printf("Salary: %g \n", salary)
	// }
	// fmt.Println("--------------------------------")

	// salary, err = gobases4.CalculateSalary(190, 1000)
	// if err != nil && errors.As(err, &errLow) {
	// 	fmt.Printf("Salary: %g, LowSalaryError: %s \n", salary, err.Error())
	// } else if err != nil && errors.As(err, &errTax) {
	// 	fmt.Printf("Salary: %g, TaxableError: %s \n", salary, err.Error())
	// } else {
	// 	fmt.Printf("Salary: %g \n", salary)
	// }

	//Archives
	// fmt.Println("---------------Read File-----------------")
	// info := gobases4.Read("/Users/arenteria/Documents/Cursos/go_from_zero/bootcamp/goBases4/customers.txt")
	// fmt.Println(info)

	// Customers := []gobases4.Customer{
	// 	gobases4.Customer{File: "AndresR.txt", Name: "Andres Rente", ID: 1, Phone: "3192735641", Home: "Here"},
	// 	gobases4.Customer{File: "AndresR.txt", Name: "Andres Rente", ID: 2, Phone: "3192735641", Home: "Here"},
	// 	gobases4.Customer{File: "AndresR.txt", Name: "Andres Rente", ID: 3, Phone: "3192735641", Home: "Here"},
	// 	gobases4.Customer{File: "AndresR.txt", Name: "Andres Rente", ID: 4, Phone: "3192735641", Home: "Here"},
	// 	gobases4.Customer{File: "AndresR.txt", Name: "Andres Rente", ID: 5, Phone: "3192735641", Home: "Here"},
	// }
	// gobases4.Write("bootcamp/goBases4/customers.txt", gobases4.Customer{ID: 6, Name: "Felipe", File: "Felipe.txt", Phone: "3192734512", Home: "Here"}, Customers)

	// fmt.Println("End of execution")

	//Channels
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go process(i, c)
	}
	fmt.Println("Program finished")
	for i := 0; i < 10; i++ {
		fmt.Println("Last Value: ", <-c)

	}
	fmt.Println("--------------------------------")

}

func process(i int, c chan int) {
	fmt.Println(i, " - Begins")
	time.Sleep((1000 * time.Millisecond))
	fmt.Println(i, " - Ends")
	c <- i
}
