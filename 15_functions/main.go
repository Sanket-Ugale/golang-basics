package main

import "fmt"

func main() {
	fmt.Println(adder(3, 4))
	fmt.Println(proAdder(3, 4, 5, 6))
}
func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
