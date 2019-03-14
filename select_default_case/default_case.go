package main

import (
	"fmt"
	"time"
)

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func main() {
	ch := make(chan string)
	go process(ch)
	
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
			case v := <- ch:
				fmt.Printf("Received value: %s\n", v)
				return
			default:
				fmt.Println("No value received yet")
		}
	}
}