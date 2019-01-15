package main

import (
	"fmt"
)

func main() {
	var a [3]int // int array with length 3
	a[0] = 25
	a[1] = 54
	a[2] = 61
	fmt.Println(a)
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)
	c := [...]string{"hello", "world"}
	fmt.Println(c)

	d := [...]string{"USA", "China", "India","Germany", "France"}
	e := d
	e[0] = "Nigeria"
	fmt.Println("a is ", d)
	fmt.Println("e is ", e)
	fmt.Println("Length is", len(e))
	for i, v := range e {
		fmt.Println(i, v)
	}
}
