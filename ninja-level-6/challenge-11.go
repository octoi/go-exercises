package main

import "fmt"

func main() {
	n := fac(10)
	fmt.Println(n)
}

func fac(n int) int {
	if n == 1 {
		return n
	} else {
		return n * fac(n - 1)
	}
}