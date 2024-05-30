package main

import (
	"fmt"
	"net/url"
)

const urlStr = "https://mindcare-app.onrender.com/api/chats"

func main() {
	fmt.Println("Welcome to Urls in golang")
	fmt.Println(urlStr)
	result, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	qparams := result.Query()
	fmt.Printf("Type is: %T", qparams)
	// fmt.Println()
	// for _, val := range qparams {
	// 	fmt.Println(val)
	// }
}
