package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Firstname string
	Lastname  string
}

func main() {

	p1 := user{
		Firstname: "Rahul",
		Lastname:  "Dravid",
	}
	p2 := user{
		Firstname: "Sachin",
		Lastname:  "Tendulkar",
	}
	fmt.Println(p1, p2)
	players := []user{p1, p2}
	fmt.Println(players)
	list, err := json.Marshal(players)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(list))

}
