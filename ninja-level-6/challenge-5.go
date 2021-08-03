package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (s square) area()  float64 {
	return s.length * s.length
}

func (c circle) area()  float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(prefix string, s shape) {
	fmt.Println(prefix, s.area())
}

func main() {
	c := circle{ radius: 12 }
	s := square{ length: 12 }

	info("Area of circle:", c)
	info("Area of square:", s)

}