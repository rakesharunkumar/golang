package main

import (
	"fmt"
)

type player struct {
	fname string
	lname string
}

type skill struct {
	player
	age     int
	batsman string
}

func main() {

	player1 := skill{
		player: player{
			fname: "Virat",
			lname: "Kohli",
		},
		batsman: "yes",
	}

	player2 := skill{
		age: 30,
		player: player{
			fname: "Rohit",
			lname: "Sharma",
		},
	}
	fmt.Println(player1.fname, player2.age)
}
