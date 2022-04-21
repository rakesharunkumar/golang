package main

import (
	"fmt"
)

//const x = 1
//const y = 3.4
//const z = "I am Happy"

const (
	x = 1
	y = 3.4
	z = "I am Happy"
)

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
}
