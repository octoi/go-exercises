package main

import (
	"fmt"
)

func main() {
	by := 1000 // birth year
	currentYear := 2021 // you can hard code this ;)

	for by < currentYear {
		fmt.Println(by)
		by++
	}
}