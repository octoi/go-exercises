package main

import "fmt"

type person struct {
	name string
}

func main() {
	p := person{ name: "James Bond" }
	changeMe(&p)
	fmt.Println(p.name)
}

func changeMe(p *person) {
	p.name = "Miss Moneypenny"
}