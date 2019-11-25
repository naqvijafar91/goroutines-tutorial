package main

import "fmt"

func write(c chan string) {
	c <- "Hello"
}

func main() {
	c := make(chan string)
	go write(c)
	fmt.Println(<-c)
}
