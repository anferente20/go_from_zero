package exercices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number int
var err error

func GetMultiplicationTable() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Write a number: ")
		if scanner.Scan() {
			number, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Wrong number.")
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(number, " * ", i, " = ", i*number)
	}

}
