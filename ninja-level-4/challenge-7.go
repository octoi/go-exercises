package main

import "fmt"

func main() {
	x := [][]string {
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}

	for id, person := range x {
		fmt.Println(id, "\n-----")
		for _, v := range person {
			fmt.Println(v)
		}
	}
}