package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is: ", diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("value 1")
	case 2:
		fmt.Println("value 2")

	case 3:
		fmt.Println("value 3")
	case 4:
		fmt.Println("value 4")
	case 5:
		fmt.Println("value 5")
	case 6:
		fmt.Println("value 6")
	default:
		fmt.Println("What was that!")
	}
}
