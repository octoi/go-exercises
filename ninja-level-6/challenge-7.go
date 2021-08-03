package main

import "fmt"

func main() {
	fn := func() {
		fmt.Println("Assigning function to a variable :)")
	}

	fn()
}