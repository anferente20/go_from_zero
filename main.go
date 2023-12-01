package main

import (
	"fmt"
	gobases3 "go_from_zero/bootcamp/goBases3"
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
	fmt.Println("--------------------------------")
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

	person := gobases3.Person{
		ID:          1,
		Name:        "Andres F",
		DateOfBirth: "05/04/1997",
	}

	employee := gobases3.Employee{
		ID:       1,
		Position: "Software Developer",
		PersonR:  person,
	}
	employee.PrintEmployee()
	fmt.Println("--------------------------------")

}
