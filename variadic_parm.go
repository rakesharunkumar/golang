package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	sum(1, 2, 3, 4, 5, 6) // series of same variable type as arguement.
}

func sum(x ...int) {
	fmt.Println(x)
	sumofvalue := 0
	fmt.Printf("%T\n", x) // prints array value of x if ... is mentioned.
	for i, v := range x {
		fmt.Println(i, v)
		sumofvalue += v
	}
	fmt.Println(sumofvalue)
}
