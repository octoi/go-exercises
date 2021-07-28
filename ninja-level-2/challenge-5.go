package main

import "fmt"

// Create a variable of type string using a raw string literal. Print it.

func main() {
	str := `You can use raw literal or back quote for strings
Which can be multiline & "" can be used in this`

	fmt.Println(str)
}