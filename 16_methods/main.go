package main

import "fmt"

func main() {
	fmt.Println("Stucts in golang")

	sanket := User{"Sanket", "sanket@gmail.com", true, 20}
	fmt.Println(sanket)
	fmt.Printf("Sanket details are: %+v\n", sanket)
	sanket.GetStatus()
	sanket.NewEmail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User active: ", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test@mail.dev"
	fmt.Println("User Email: ", u.Email)
}
