package main

import "fmt"

func main() {
	v := foo()()
	fmt.Println(v)
}

func foo() func() string {
	return func() string {
		return "Somewhere in the function stack"
	}
}