package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type triangle struct {
	baseLength float64
	height     float64
}

type square struct {
	sideLength float64
}

func main() {
	t := triangle{baseLength: 4, height: 2}
	fmt.Printf("Triangle area is %.2f\n", t.area())

	s := square{sideLength: 2}
	fmt.Printf("Square area is %.2f\n", s.area())
}

func (t triangle) area() float64 {
	return t.baseLength * t.height / 2
}

func (s square) area() float64 {
	return math.Pow(s.sideLength, 2)
}
