package main

import (
	"fmt"
)

type person struct {
	firstname string
	lastname  string
}

type dob struct {
	person
	year  int
	month int
}

func main() {
	person_details := dob{
		person: person{
			"Sourav",
			"Ganguly",
		},
		year: 1980,
	}
	// fmt.Println(person_details)
	// fmt.Printf("Birth year of %s is %d\n", person_details.firstname, person_details.year)
	person_details.call_method()
}

func (p person) call_method() {
	fmt.Println(p.firstname)
	//fmt.Println(p.age) only one recever is possible for one func.
}
