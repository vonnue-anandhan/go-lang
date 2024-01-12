package main

import (
	"fmt"
	"math"
)

// **************** Interface ****************

type shape interface { // Groups types together based on their methods
	area() float64
	circumf() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

// Square methods

func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumf() float64 {
	return s.length * 4
}

// Circle methods

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s shape) {
	fmt.Printf("Area of %T is : %0.2f \n", s, s.area())
	fmt.Printf("Circumference of %T is : %0.2f \n", s, s.circumf())
}
