package main

import "fmt"

func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

func main() {
	price, no := 40, 5
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is ", totalPrice)
}
