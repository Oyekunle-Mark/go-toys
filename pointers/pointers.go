package main

import (
	"fmt"
)

func alter(val *int) {
	*val = 21
}

func main() {
	b := 25
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)

	//to get the value a is pointing at
	fmt.Printf("value of b is %d\n", *a)

	//mathematical operations can be performed on a destructured pointer
	x := 100
	y := &x
	*y++
	fmt.Println("Value at one increased one is", x)

	//pointers can be passed as arguments to a function
	alter(y)
	fmt.Println("alter valued of x is:", x)
}
