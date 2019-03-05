package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("Hello goroutine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("Hello goroutine awake to write done")
	done <- true
}

func main() {
	done := make(chan bool)
	fmt.Println("Main goroutine to call hello goroutine")
	go hello(done)
	<- done
	fmt.Println("main received data")
}