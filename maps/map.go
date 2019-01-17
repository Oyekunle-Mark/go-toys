package main

import (
	"fmt"
)

func main() {
	//creating maps
	employee := make(map[string]int)
	employee["salary"] = 1230
	employee["age"] = 18
	fmt.Println("Employee map contents:", employee)

	//creating and initializing maps
	person := map[string]string{
		"name": "John Doe",
		"address": "10, Lagos Drive.",
		"sex": "male",
	}
	person["nationality"] = "Nigerian"
	fmt.Println("Person map contents:", person)

	//accessing items of map
	salary := "salary"
	fmt.Println("Employee", salary, "is", employee[salary])

	// to know if a key is present in a map
	value, ok := person["sex"]
	if ok {
		fmt.Println("Sex is recorded as:", value)
	} else {
		fmt.Println("Sex is not recorded")
	}

	//to loop over all the key and value pair in the map use for range
	fmt.Println("The items in the person map are:")
	for key, value := range person {
		fmt.Printf("person[%s] = %s\n", key, value)
	}

	//use the delete function to delete from the map
	delete(employee, "salary")
	fmt.Println("The new contents of the employee map are:", employee)

	//length with the function len
	fmt.Printf("Length of the person map is %d\n", len(person))
}
