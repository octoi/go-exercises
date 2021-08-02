package main

import "fmt"

func main() {
	s := struct {
		first string
		friends map[string] int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"bond": 1,
			"todd": 2,
			"me": 3,
		},
		favDrinks: []string{"pepsi", "water"},
	}

	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)
}