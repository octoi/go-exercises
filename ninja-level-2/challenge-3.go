package main

import "fmt"

// Create TYPED and UNTYPED constants. Print the values of the constants.

const (
	a = 42 // UNTYPED
	b int = 42 // TYPED
)

func main() {
	fmt.Println(a, b)
}