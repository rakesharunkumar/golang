package main

import (
	"fmt"
)

type owntype int

var x owntype
var y owntype

func main() {
	fmt.Println(x)
	fmt.Printf("%T", x)
	x = 42
	y = owntype(42)
	fmt.Println(y)
	fmt.Printf("%T", y)

}
