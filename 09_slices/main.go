package main

import "fmt"

func main() {
	fmt.Println("Welcome to slices")
	var fruitList = []string{"Apple", "Tomato"}
	fmt.Printf("Type of furitList is: %T\n", fruitList)
	fruitList = append(fruitList, "Banana")
	fmt.Println(fruitList)
	fmt.Println("", fruitList[1:])
	fmt.Println("", fruitList[1:2])
	fmt.Println("", fruitList[:2])
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 345
	highScores[2] = 645
	highScores[3] = 867
	// highScores[4]=777
	fmt.Println(highScores)
	highScores = append(highScores, 333, 444, 555)
	fmt.Println(highScores)

	var courses = []string{"reacjs", "javascript", "swift", "python"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
