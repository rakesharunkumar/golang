package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	Firstname string
	Lastname  string
	Age       int
	Quote     []string
}

func main() {

	p1 := user{
		Firstname: "Rahul",
		Lastname:  "Dravid",
		Age:       40,
		Quote: []string{"Never loose the bat",
			"Always leave the hard ball"},
	}
	p2 := user{
		Firstname: "Sachin",
		Lastname:  "Tendulkar",
		Quote: []string{"Hit 4 for yorker",
			"Straight drive for spinners"},
	}

	players := []user{p1, p2}
	fmt.Println(players)
	err := json.NewEncoder(os.Stdout).Encode(players)
	if err != nil {
		fmt.Println("Something wrong", err)
	}

}
