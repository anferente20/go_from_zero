package gobases1

import (
	"bufio"
	"fmt"
	"os"
)

var label string

func GetLetters() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Write a word: ")
	if scanner.Scan() {
		label = scanner.Text()
	}

	for i := 0; i < len(label); i++ {
		fmt.Print(string(label[i]), " ")
	}
	fmt.Println()
}
