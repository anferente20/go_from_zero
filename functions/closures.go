package functions

import "fmt"

func table(value int) func() int {
	secuence := 0
	return func() int {
		secuence++
		return value * secuence
	}
}

func CallClosure(value int) {
	MyTable := table(value)
	for i := 1; i <= 10; i++ {
		fmt.Println(value, " * ", i, " = ", MyTable())
	}
}
