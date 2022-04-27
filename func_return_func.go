package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	x := return_func()
	fmt.Println(x) //Prints the Type of x
	i := x()
	fmt.Println(i)
}

func return_func() func() string {
	return func() string {
		y := "funcwithinfunc"
		return y
	}
}
