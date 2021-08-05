package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int, 1)

	// using func literal, aka, anonymous self-executing func
	go func() {
		c1 <- 42
	}()

	c2 <- 42 // a buffered channel

	fmt.Println(<-c1)
	fmt.Println(<-c2)
}