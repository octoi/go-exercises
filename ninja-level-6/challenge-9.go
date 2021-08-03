package main

import "fmt"

func main() {
	foo(func(s string) {
		fmt.Println(s)
	})
}

func foo(fn func(s string)) {
	fn("Callback")
}