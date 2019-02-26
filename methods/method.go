package main

import (
	"fmt"
)

type Employee struct {
	name string
	salary int 
	currency string
}

// method with Employee receiver type
func (e Employee) displayEmployee() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func main() {
	emp := Employee {
		name: "John Doe",
		salary: 5000,
		currency: "$",
	}

	emp.displayEmployee();
}