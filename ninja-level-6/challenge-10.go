package main

import "fmt"

func main() {
	fn := foo()

	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())

	// re assigning
	fn = foo()

	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
}

func foo() func() int {
	x := 0

	return func() int {
		x++
		return x
	}
}