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

	m := map[string] person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, p := range m {
		fmt.Println(p.first)
		fmt.Println(p.last)

		for i, v := range p.favFlavours {
			fmt.Println(i, v)
		}
	}

}
