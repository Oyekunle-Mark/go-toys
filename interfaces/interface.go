package main

import (
	"fmt"
)

type SalaryCalculator interface {
	CalculateSalary() int
}

type Describer interface {
	Describe()
}

type Permanent struct {
	empId int
	basicPay int
	pf int
}

type Contract struct {
	tmpId int
	basicPay int
}

type Person struct {
	name string
	age int
}

func (p Permanent) CalculateSalary() int {
	return p.basicPay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicPay
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total expense per month $%d\n", expense)
}

func (p Person) Describe() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType(i interface{}) {
	switch v := i.(type) {
		case Describer:
			v.Describe()
		default:
			fmt.Printf("Unknown type\n")
	}
}

func main() {
	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 20}
	cemp1 := Contract{3, 3000}

	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)

	findType("Oyekunle")
	p := Person{
		name: "Oyekunle Mark",
		age: 26,
	}
	findType(p)
}