package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("Go is awesome ;)")
}

func foo() {
	defer func() {
		fmt.Println("running anonymous function")
	}()
	fmt.Println("foo ran")
}
