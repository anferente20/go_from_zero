package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Name string
var State bool
var Salary float32
var Date time.Time

func OtherVariables() {
	Name = "Andrés"
	State = true
	Salary = 1997.0504
	Date = time.Now()
	fmt.Println("--------------------------------")
	fmt.Println("----------- Varibales -----------")
	fmt.Println("Name = ", Name)
	fmt.Println("State = ", State)
	fmt.Println("Salary = ", Salary)
	fmt.Println("Date = ", Date)
	fmt.Println("--------------------------------")
}

func CastToText(number int) (bool, string) {
	text := strconv.Itoa(number)
	return true, text
}
