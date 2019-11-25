package main

import (
	"fmt"
	"sync"
)

func hello(wg *sync.WaitGroup) {
	fmt.Println("Hello world goroutine")
	wg.Done()
}
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go hello(wg)
	wg.Wait()
	fmt.Println("main goroutine")
}
