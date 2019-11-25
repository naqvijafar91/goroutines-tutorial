package main

import (
	"fmt"
)

func hello(ch chan string) {
	ch <- "Hello world goroutine"
}
func main() {
	ch := make(chan string)
	go hello(ch)
	responseFromGoroutine := <-ch
	fmt.Println(responseFromGoroutine)
	fmt.Println("main goroutine")
}
