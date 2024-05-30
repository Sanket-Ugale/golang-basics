package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golan")
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Bunana"

	fmt.Println("Fruit list is: ", fruitList)

	var vegList = [5]string{"potato", "beans", "mashroom"}
	fmt.Println("vegy list is: ", len(vegList))
}
