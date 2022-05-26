package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func main() {
	sq := square{
		sideLength: 10,
	}

	tr := triangle{
		base:   6,
		height: 3,
	}

	fmt.Println("Square below")
	printArea(sq)

	fmt.Println("Triangle below")
	printArea(tr)
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func (s square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}

func printArea(s shape) {
	fmt.Println("Area of shape is", s.getArea())
}
