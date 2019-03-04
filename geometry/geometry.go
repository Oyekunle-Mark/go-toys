package main

import (
	"fmt"
	"github.com/oyekunle-mark/go-toys/geometry/rectangle"
	"log"
)

var rectLen, rectWidth float64 = 6, 21 

func init() {
	fmt.Println("main package initialised")
	if (rectLen < 0) {
		log.Fatal("length is less than zero")
	}

	if (rectWidth < 0) {
		log.Fatal("width is less than zero")
	}
}

func main() {
	fmt.Println("Geometrical shape properties")
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	fmt.Printf("diagonal of the rectangle %.2f\n", rectangle.Diagonal(rectLen, rectWidth))
}
