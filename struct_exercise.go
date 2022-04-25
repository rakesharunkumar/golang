package main

import (
	"fmt"
)

type person struct {
	firstname     string
	lastname      string
	fav_ice_cream []string
}

func main() {
	p1 := person{
		firstname:     "Rohit ",
		lastname:      "Sharma",
		fav_ice_cream: []string{"Vanilla", "Strawberry"},
	}
	p2 := person{
		firstname:     "Sachin ",
		lastname:      "Tendulkar",
		fav_ice_cream: []string{"Butterscotch", "Black Current"},
	}
	fmt.Println(p1)
	for i, v := range p1.fav_ice_cream {
		fmt.Println(i, v)
	}
	fmt.Println(p2)
	for i, v := range p2.fav_ice_cream {
		fmt.Println(i, v)
	}
}
