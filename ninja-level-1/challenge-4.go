package main

import "fmt"

type hotdog int

func main() {
	var x hotdog

	fmt.Println(x)
	fmt.Printf("%T\n", x) // %T will print the type

	x = 42
	fmt.Println(x)
}