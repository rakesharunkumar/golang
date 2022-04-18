package main

import (
	"fmt"
)

var x int = 10

type mytype int

var y mytype = 11

func main() {
	fmt.Println(x)
	fmt.Printf("type value of x is: %T", x)
	fmt.Println(y)
	fmt.Printf("type value of y is: %T", y)
}
