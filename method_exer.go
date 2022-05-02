package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	person1 := person{
		"Anil",
		"Kumble",
		40,
	}
	//fmt.Println(person1)
	person1.speak()
}

func (p person) speak() {
	fmt.Println(p.first)
	fmt.Println(p.age)
}
