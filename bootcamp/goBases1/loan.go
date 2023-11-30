package gobases1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var age int
var hasJob int
var salary int
var hasAntiqity int
var err error

func CanBorrow() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Write age: ")
		if scanner.Scan() {
			age, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Wrong number.")
			} else {
				break
			}
		}
	}

	for {
		fmt.Println("Has job? (1-true, 2- false): ")
		if scanner.Scan() {
			hasJob, err = strconv.Atoi(scanner.Text())
			if err != nil || hasJob < 1 || hasJob > 2 {
				fmt.Println("Wrong number.")
			} else {
				break
			}
		}
	}

	for {
		fmt.Println("Write salary: ")
		if scanner.Scan() {
			salary, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Wrong number.")
			} else {
				break
			}
		}
	}
	for {
		fmt.Println("Has antiquity? (1-true, 2- false): ")
		if scanner.Scan() {
			hasAntiqity, err = strconv.Atoi(scanner.Text())
			if err != nil || hasAntiqity < 1 || hasAntiqity > 2 {
				fmt.Println("Wrong number.")
			} else {
				break
			}
		}
	}

	if age > 22 && hasJob == 1 && hasAntiqity == 1 {
		if salary > 1000 {
			fmt.Println("You can borrow us money, with no interest")
		} else {
			fmt.Println("You can borrow us money, with interest")

		}
	} else {
		fmt.Println("You can't borrow us money")

	}

}
