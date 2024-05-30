package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Go Channels")
	wg := &sync.WaitGroup{}
	myCh := make(chan int)
	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		myCh <- 5
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}
