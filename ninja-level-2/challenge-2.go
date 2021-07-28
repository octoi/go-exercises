package main

import "fmt"

// Using the following operators, write expressions and assign their values to variables:
// 	a ==
// 	b <=
// 	c >=
//  d !=
// 	e <
//  f >
// Now print each of the variables.


func main() {
	num1 := 42 // does't have to be 42
	num2 := 43 // does't have to be 43

	a := num1 == num2
	b := num1 <= num2
	c := num1 >= num2
	d := num1 != num2
	e := num1 < num2
	f := num1 > num2

	fmt.Println(a, b, c, d, e, f)
}