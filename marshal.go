package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Fname string
	Lname string
	Age   int
}

func main() {
	p1 := person{
		Fname: "Sachin",
		Lname: "Tendulkar",
		Age:   40,
	}
	p2 := person{
		Fname: "Veerendar",
		Lname: "Sehwag",
		Age:   39,
	}

	players := []person{p1, p2}

	fmt.Println(players)

	values, err := json.Marshal(players)
	if err != nil { //Error catching
		fmt.Println(err)
	}
	fmt.Println(values)         // This will print numbers now convert it into string and change the variable starting letter to capital
	fmt.Println(string(values)) // This will print the json structure

}
