package main

import (
	"fmt"
	exercices "go_from_zero/exercises"
	"go_from_zero/variables"
)

func main() {

	fmt.Println("Hello world")

	variables.ShowIntegers()
	variables.OtherVariables()

	fmt.Println("--------------------------------")
	fmt.Println("----------- Cast Integer -----------")
	var state, text = variables.CastToText(190919)
	fmt.Println("State = ", state)
	fmt.Println("Text = ", text)
	fmt.Println("--------------------------------")

	fmt.Println("--------------------------------")
	fmt.Println("----------- EX 1 -----------")
	fmt.Println(exercices.IsMoreThan100("a "))
	fmt.Println("--------------------------------")

}
