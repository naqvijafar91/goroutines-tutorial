package main

import (
	"fmt"
	"sync"
)

var num = 0

func increment(wg *sync.WaitGroup, mut *sync.Mutex) {
	mut.Lock()
	num = num + 1
	mut.Unlock()
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go increment(wg, mut)
	}
	wg.Wait()
	fmt.Println(num)
}
