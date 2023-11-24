package variables

import "fmt"

func ShowIntegers() {
	var commonInteger int
	integer32 := int32(10)
	integer64 := int64(100)
	fmt.Println("--------------------------------")
	fmt.Println("----------- Integers -----------")
	fmt.Println("Common Integer = ", commonInteger)
	fmt.Println("32-bit Integer = ", integer32)
	fmt.Println("64-bit Integer = ", integer64)
	fmt.Println("--------------------------------")

}
