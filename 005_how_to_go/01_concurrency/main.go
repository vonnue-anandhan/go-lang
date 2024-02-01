package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius int
}

func (c circle) area() int {
	return 2 * c.radius * c.radius
}

type shape interface {
	area() int
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{radius: 5}
	info(c)
}
