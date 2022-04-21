package main

import (
	"fmt"
)

func main() {

	myvariable := "testswitch"
	switch myvariable {
	case "notmyvariable":
		fmt.Println("Not the testswitch")
	case "testswitch":
		fmt.Println("Myvariable testswitch")

	}
}
