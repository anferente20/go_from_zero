package gobases1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number int
var repeat int

func GetSalary() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Write your employee salary: ")
		if scanner.Scan() {
			number, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Wrong number.")
			} else {
				salary := float32(number)
				switch {
				case number < 50000:
					fmt.Println("Your employee salary is: ", salary)

				case number > 50000 && number <= 150000:
					fmt.Println("Your employee salary is: ", salary*0.83)
				case number > 150000:
					fmt.Println("Your employee salary is: ", salary*0.73)
				default:
					fmt.Println("Your employee salary is: ", salary)
				}
			}
		}
		fmt.Println("do you wanna get another salary? (1 - yes): ")
		if scanner.Scan() {
			repeat, err = strconv.Atoi(scanner.Text())
			if err != nil || repeat != 1 {
				fmt.Println("Thanks.")
				return

			}
		}
	}
}
