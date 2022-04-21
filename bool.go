package main

import (
	"fmt"
)

var x bool
var y = 20
var z = 30

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)
	fmt.Println(y == z)
}
