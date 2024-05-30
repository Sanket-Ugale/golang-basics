package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("index is %v and value is %v \n", index, day)
	}
	roughValue := 1

	for roughValue < 10 {
		if roughValue == 2 {
			goto lco
		}
		if roughValue == 5 {
			// break
			roughValue++
			continue
		}
		fmt.Println("value is: ", roughValue)
		roughValue++
	}
lco:
	fmt.Println("Jumped")
}
