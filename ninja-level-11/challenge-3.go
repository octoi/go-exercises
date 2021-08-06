package main

import "fmt"

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("error: %v", ce.info)
}

func main() {
	c1 := customErr{info: "Never gonna give yo up, never gonna let you down !!"}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}