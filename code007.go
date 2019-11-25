package main

import (
	"fmt"
	"sync"
)

var num = 0

func increment(wg *sync.WaitGroup) {
	num = num + 1
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	for i:=0;i<500;i++ {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println(num)
}
