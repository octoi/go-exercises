package main

import "fmt"

// Write a program that
// 	- assigns an int to a variable
// 	- prints that int in decimal, binary, and hex
// 	- shifts the bits of that int over 1 position to the left, and assigns that to a variable
// 	- prints that variable in decimal, binary, and hex



func main() {
	number := 42

	// %d = decimal
	// %b = binary
	// %x = hex, %#x = to add `0x` before the hex value
	// \t = tab space
	// \n = new line

	fmt.Printf("%d\t%b\t%#x", number, number, number)

	number = number << 1
	fmt.Printf("\n%d\t%b\t%#x", number, number, number)
}
