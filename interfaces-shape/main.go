package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
	getType() string
}

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (square) getType() string {
	return "Square"
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func (triangle) getType() string {
	return "Triangle"
}

func printArea(s shape) {
	fmt.Println("Area of", s.getType(), "is:", s.getArea())
}

func main() {
	square := square{sideLength: 5}
	triangle := triangle{base: 5, height: 5}

	printArea(square)
	printArea(triangle)
}
