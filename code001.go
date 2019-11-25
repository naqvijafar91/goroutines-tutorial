package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}
func main() {
	go hello()
	// If you comment the below line, nothing will be printed
	// as the main goroutine will exit before the other has executed
	time.Sleep(1 * time.Second)
	fmt.Println("main goroutine")
}
