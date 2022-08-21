package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Printf("The area is %v\n", s.getArea())
}

func main() {

	sq := square{sideLength: 10}
	tr := triangle{height: 10, base: 5}

	printArea(sq)
	printArea(tr)
}
