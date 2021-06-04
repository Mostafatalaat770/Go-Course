package main

import (
	"fmt"
	"math"
)

type square struct {
	width  float64
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}
func main() {
	s := square{
		length: 10,
		width:  5.5,
	}
	c := circle{
		radius: 10,
	}
	info(s)
	info(c)
}
