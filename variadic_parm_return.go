package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	sumofvalues := sum(1, 2, 3, 4, 5, 6) // series of same variable type as arguement.
	fmt.Println(sumofvalues)
}

func sum(x ...int) int {
	fmt.Println(x)

	fmt.Printf("%T\n", x) // prints array value of x if ... is mentioned.
	sumofvalue := 0
	for i, v := range x {
		fmt.Println(i, v)
		sumofvalue += v
		fmt.Printf("Sum of %d with %d is %d\n", v, sumofvalue, sumofvalue)

	}
	return sumofvalue
}
