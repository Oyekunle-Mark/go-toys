package employee

import "fmt"

type employee struct {
	firstname string
	lastname string
	totalLeaves int
	leavesTaken int
}

func New(firstname string, lastname string, totalLeaves int, leavesTaken int) employee {
	e := employee {firstname, lastname, totalLeaves, leavesTaken}
	return e
}

func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.firstname, e.lastname, (e.totalLeaves - e.leavesTaken))
}