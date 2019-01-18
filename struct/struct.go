package main

import (
	"fmt"
)

//creating a named structure
type Address struct {
	city, state string
}

type Person struct {
	name string
	age int
	address Address
}

type Employee struct {
	firstName, lastName string
	age, salary	int
}

func main() {
	//creating structures using field names
	emp1 := Employee{
		firstName: "john",
		age: 29,
		salary: 1000,
		lastName: "doe",
	}

	emp2 := Employee{"jane", "doe", 21, 2000}

	fmt.Println("Employee 1 struct details are", emp1)
	fmt.Println("Employee 2 struct details are", emp2)

	//creating anonymous structures
	emp3 := struct {
		firstName, lastName string
		age, salary int
	}{
		firstName: "bola",
		lastName: "sumbo",
		age: 52,
		salary: 5000,
	}

	fmt.Println("Employee 3 struct details are",emp3)

	//an uninitialized struct has zero values
	var emp4 Employee
	fmt.Println("Unitialized struct", emp4)

	emp4.firstName = "kunle"
	emp4.lastName = "mark"
	fmt.Println("Employee 4", emp4)

	//the individual fields in a struct can be accessed using dot notation
	emp5 := Employee {
		firstName: "john",
		lastName: "wick",
	}
	fmt.Println("Employee 5 FirstName", emp5.firstName)
	fmt.Println("Employee 5 lastName", emp5.lastName)

	//pointers can also be used with structs
	emp6 := &Employee {"ladies", "gentlemen", 1, 2}
	fmt.Println("Details of the pointer to the struct", emp6.firstName)

	//nested structs
	var person Person 
	person.name = "halo"
	person.age = 15
	person.address = Address {
		city: "lagos",
		state: "lagos",
	}
	fmt.Println("Person", person)
}
