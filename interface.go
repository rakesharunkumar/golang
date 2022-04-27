package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}

type person_details struct {
	person
	doy   int
	month string
	day   string
}

type players interface {
	batsman()
}

func main() {
	player1 := person_details{
		person: person{
			"Sachin",
			"Tendulkar",
		},
		doy:   1975,
		month: "April",
	}
	player2 := person_details{
		person: person{
			"Rahul",
			"Dravid",
		},
		doy:   1978,
		month: "August",
	}
	player3 := person{
		"Sourav",
		"Ganguly",
	}
	fmt.Println(player2)
	player1.batsman()
	player2.batsman()
	players_list(player1)
	players_list(player2)
	players_list(player3)
}

func (p1 person) batsman() {
	fmt.Println(p1.fname)
}

func players_list(p players) {
	fmt.Println("Printing from Interface", p)
}
