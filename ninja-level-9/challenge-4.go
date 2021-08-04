package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	incrementer := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			iv := incrementer
			iv++
			incrementer = iv
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(incrementer)

}