package main

import "fmt"

func main() {
	n := foo()
	num, str := bar()

	fmt.Println(n)
	fmt.Println(str, "loves the number", num)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 42, "James bond"
}