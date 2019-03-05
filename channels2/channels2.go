package main

import (
	"fmt"
)

func digits(number int, dchn chan int) {
	for number != 0 {
		digit := number % 10
		dchn <- digit
		number /= 10
	}
	close(dchn)
}

func calcSquares(number int, sqop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	sqop <- sum
}

func calcCubes(number int, cbop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cbop <- sum
}

func main() {
	number := 589
	sqop := make(chan int)
	cbop := make(chan int)
	go calcCubes(number, cbop)
	go calcSquares(number, sqop)
	squares, cubes := <-sqop, <-cbop
	fmt.Println("Final output", squares+cubes) 
}