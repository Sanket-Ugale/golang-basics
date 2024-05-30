package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps")
	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println(languages)
	fmt.Println("JS shorts for:", languages["JS"])
	delete(languages, "RB")
	fmt.Println(languages)

	// loops in golang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
