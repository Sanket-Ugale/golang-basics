package main

import (
	"fmt"
	"net/http"

	"github.com/sanket-ugale/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB apis")
	r := router.Router()
	fmt.Println("Server is getting started ......")
	http.ListenAndServe("0.0.0.0:80", r)
	fmt.Println("Listening at port 80 .........")
}
