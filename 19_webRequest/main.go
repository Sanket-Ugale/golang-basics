package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://mindcare-app.onrender.com"

func main() {
	fmt.Println("First Web Request")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response type is %T", response)

	databyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(databyte))
	response.Body.Close()
}
