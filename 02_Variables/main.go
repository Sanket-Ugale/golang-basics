package main

import (
	"fmt"
)

func main() {
	var name string = "Gopher"
	fmt.Println(name)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)
	fmt.Println()

	var age int = 20
	fmt.Println(age)
	fmt.Printf("Variable is of type: %T\n", age)
	fmt.Println()

	var data1 uint8 = 233
	fmt.Println(data1)
	fmt.Printf("Variable is of type: %T\n", data1)
	fmt.Println()

	var data2 uint16 = 12345
	fmt.Println(data2)
	fmt.Printf("Variable is of type: %T\n", data2)
	fmt.Println()

	var data3 uint32 = 1234567899
	fmt.Println(data3)
	fmt.Printf("Variable is of type: %T\n", data3)
	fmt.Println()

	var data4 uint64 = 9999999999999999999
	fmt.Println(data4)
	fmt.Printf("Variable is of type: %T\n", data4)
	fmt.Println()
}
