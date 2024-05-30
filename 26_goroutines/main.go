package main

import (
	"fmt"
	"net/http"
	"sync"
)

// func main() {
// 	// fmt.Println("Goroutine")
// 	go greeter("Hii")
// 	greeter("Hello")
// }

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

var signals = []string{"test"}
var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	websitelist := []string{
		"https://google.com",
		"https://linkedin.com",
		"https://github.com",
		"https://sanketugale.netlify.app",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Endpoint Error")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
