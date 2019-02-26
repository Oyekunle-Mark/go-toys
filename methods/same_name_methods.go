package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	length int
	width int
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius
}

func main() {
	r := Rectangle {
		length: 10,
		width: 5,
	}

	c := Circle {
		radius: 7.5,
	}

	fmt.Printf("Area of the rectangle is %d\n", r.Area())
	fmt.Printf("Area of the circle is %f\n", c.Area())
}