package main

import (
	"fmt"
    "equal/equal_map"
)

func main() {
	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}

	map2 := map[string]int{
		"one": 5,
		"two": 3,
	}
    
	state := equal.Compare(map1, map2)
	if state {
        fmt.Println("Both maps are equal")
    } else {
        fmt.Println("Maps not equal")
    }
}
