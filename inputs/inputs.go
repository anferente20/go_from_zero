package inputs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number1 int
var number2 int
var label string
var err error

func GetNumbers() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Set number 1: ")
	if scanner.Scan() {
		number1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Wrong information setted. " + err.Error())
		}
	}

	fmt.Println("Set number 2: ")
	if scanner.Scan() {
		number2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Wrong information setted. " + err.Error())
		}
	}

	fmt.Println("Set label: ")
	if scanner.Scan() {
		label = scanner.Text()
	}
	fmt.Println(label, number1*number2)
}
