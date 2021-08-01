package main

import "fmt"

func main() {
	by := 1000 // your birth year
	currentYear := 2021 // current year

	for {
		if by == currentYear {
			break
		}

		fmt.Println(by)
		by++
	}
}
