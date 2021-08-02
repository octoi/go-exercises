package main

import "fmt"

type person struct {
	first string
	last string
	favFlavours []string
}

func main() {
	p1 := person{
		first: "James",
		last: "Bond",
		favFlavours: []string{
			"vanilla", "chocolate", "martini",
		},
	}

	p2 := person{
		first: "Miss",
		last: "Moneypenny",
		favFlavours: []string{
			"strawberry", "chocolate", "vanilla",
		},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)

	for i, v := range p1.favFlavours {
		fmt.Println(i + 1, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)

	for i, v := range p2.favFlavours {
		fmt.Println(i + 1, v)
	}

}
