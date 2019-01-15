package main

import "fmt"

func countries() []string {
	countries := []string{"USA", "Nigeria", "Kenya", "Madagascar"}
	neededCountries := countries[:len(countries)-2]
	countriesCopy := make([]string, len(neededCountries))
	copy(countriesCopy, neededCountries)
	return countriesCopy
}

func main() {
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}
