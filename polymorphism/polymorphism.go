package main

import "fmt"

type Income interface {
	calculate() int
	source() string
}

type fixedBilling struct {
	projectName string
	billedAmount int
}

type timeAndMaterial struct {
	projectName string
	noOfHours int
	hourlyRate int
}

func (fb fixedBilling) calculate() int {
	return fb.billedAmount
}

func (fb fixedBilling) source() string {
	return fb.projectName
}

func (tm timeAndMaterial) calculate() int {
	return tm.noOfHours + tm.hourlyRate
}

func (tm timeAndMaterial) source() string {
	return tm.projectName
}

func calculateIncome(inc []Income) {
	netIncome := 0
	for _, income := range inc {
		fmt.Printf("Income from %s = %d\n", income.source(), income.calculate())
		netIncome += income.calculate()
	}

	fmt.Printf("Net income of firm = %d", netIncome)
}

func main() {
	project1 := fixedBilling{"Project 1", 2000}
	project2 := fixedBilling{"Project 2", 3000}
	project3 := timeAndMaterial{"Project 3", 100, 30}

	incomeStreams := []Income{project1, project2, project3}
	calculateIncome(incomeStreams)
}