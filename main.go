package main

import (
	"fmt"
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

}
