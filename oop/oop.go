package main

import "github.com/oyekunle-mark/go-toys/oop/employee"

func main()  {
	e := employee.New("John", "Doe", 40, 20)
	e.LeavesRemaining()
}