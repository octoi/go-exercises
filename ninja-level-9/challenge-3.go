package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incrementer := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			iv := incrementer
			runtime.Gosched()
			iv++
			incrementer = iv
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(incrementer)

}