package main

import (
	"fmt"
)

func main() {
	 x := "James Bond"

	 if x == "Moneypenny" {
	 	fmt.Println(x, "is Moneypenny")
	 } else if x == "James Bond" {
	 	fmt.Println(x, "is James Bond")
	 } else {
		 fmt.Println(x, "is something")
	 }
}
