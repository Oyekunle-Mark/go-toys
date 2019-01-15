package main

import (
	"fmt"
)

func main() {
	cars := []string{"Ferari", "Danfo", "Molue", "Marwa"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))
	cars = append(cars, "Lexus", "Benz", "Ford", "Range")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars))

	veggies := []string{"potatoes", "efo", "ewedu", "ugwu"}
	fruits := []string{"shombo", "shoko"}
	nVeggies := append(veggies, "Suya")
	food := append(nVeggies, fruits...)
	fmt.Println("food:",  food)
}
