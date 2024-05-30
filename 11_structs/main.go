package main

import "fmt"

func main() {
	fmt.Println("Stucts in golang")

	sanket := User{"Sanket", "sanket@gmail.com", true, 20}
	fmt.Println(sanket)
	fmt.Printf("Sanket details are: %+v\n", sanket)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
