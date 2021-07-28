package main

import "fmt"

// Using iota, create 4 constants for the NEXT 4 years.
// Print the constant values.

const (
	a = 2021 + iota
	b = 2021 + iota
	c = 2021 + iota
	d = 2021 + iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}