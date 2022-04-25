package main

import (
	"fmt"
)

type person struct {
	firstname string
	lastname  string
}

func main() {
	p1 := person{
		firstname: "Rohit",
		lastname:  "Sharma",
	}
	p2 := person{
		firstname: "Virat",
		lastname:  "Kohli",
	}
	fmt.Println(p1)
	fmt.Println(p2.lastname)
}
