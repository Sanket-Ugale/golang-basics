package main

import "fmt"

func main() {
	fmt.Println("If Else in golang")
	loginCount := 23
	var result string
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "watch out"
	} else {
		result = "Exactly 10 login count"
	}
	fmt.Println(result)
}
