package main

import "fmt"

func main() {
	defer fmt.Println("Three")
	defer fmt.Println("Two")
	defer fmt.Println("One")
	fmt.Println("Hello")
	myDefer()
	var div int = 207 / 2
	div = div - 4
	fmt.Println(div)
}
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
