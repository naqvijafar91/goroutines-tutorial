package main

import (
	"fmt"
	"time"
)

func task1(ch1 chan string) {
	time.Sleep(5 * time.Second)
	ch1 <- "Task 1 Complete"
}

func task2(ch2 chan string) {
	time.Sleep(2 * time.Second)
	ch2 <- "Task 2 Complete"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go task1(ch1)
	go task2(ch2)
	select {
	case str1 := <-ch1:
		fmt.Println(str1)
	case str2 := <-ch2:
		fmt.Println(str2)
	}

}
