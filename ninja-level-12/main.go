package main

import (
	"fmt"
	"ninja-level/ninja-level-12/dog" // this might be fail to work depends upon your GOROOT
)

type canine struct {
	name string
	age int
}

func main() {
	fido := canine{
		name: "Fido",
		age: dog.Years(10),
	}

	fmt.Println(fido)
}
