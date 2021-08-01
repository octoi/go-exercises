package main

import "fmt"

func main() {
	switch {
		case false:
			fmt.Println("it will never print")
		case true:
			fmt.Println("this will print")
	}
}
