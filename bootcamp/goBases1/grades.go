package gobases1

import (
	"fmt"
)

func CalculateGrade(grades ...int) {
	var resultado int
	for _, value := range grades {
		resultado += value
	}
	fmt.Println(grades)
	fmt.Println("Your grade is: ", resultado/len(grades))
}
