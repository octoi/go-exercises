package main

import "fmt"

// Write a program that prints a number in decimal, binary, and hex

func main() {
	const number int = 42 // does't have to `42` ;)

	// %d = decimal
	// %b = binary
	// %x = hex, %#x = to add `0x` before the hex value

	fmt.Printf("Decimal: %d\n", number)
	fmt.Printf("Binary: %b\n", number)
	fmt.Printf("Hex: %#x\n", number)
}