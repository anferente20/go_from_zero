package functions

import (
	"fmt"
)

func Calcules(number1 int, number2 int) {
	sum := func(num1 int, num2 int) int {
		return num1 + num2
	}

	fmt.Println(number1, " * ", number2, " = ", sum(number1, number2))

}
