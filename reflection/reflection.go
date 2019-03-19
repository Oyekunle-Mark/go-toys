package main

import (
	"fmt"
	"reflect"
)

type order struct {
	orderId int
	customerId int
}

type employee struct {
	name string
	id int
	address string
	salary int
	country string
}

func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("Insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch  v.Field(i).Kind() {
				case reflect.Int:
					if i == 0 {
						query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
					} else {
						query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
					}
				case reflect.String:
					if i == 0 {
						query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
					} else {
						query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
					}
				default:
					fmt.Println("Unsurpported type")
					return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return
	}
	fmt.Println("Unsupported type")
}

func main() {
	o := order {
		orderId: 456,
		customerId: 56,
	}
	createQuery(o)
 
	e := employee {
		name: "Oyekunle",
		id: 666,
		address: "Lagos",
		salary: 120000,
		country: "Nigeria",
	}
	createQuery(e)

	i := 90
	createQuery(i)
}