package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Println("Value of x before", x)
	fmt.Println("Address of x before", &x)
	func1(&x)
	fmt.Println("Value of x after", x)
	fmt.Println("Address of x after", &x)

	fmt.Println(x)
}

func func1(y *int) {
	fmt.Println("Value of y", y)
	fmt.Println("Address of y", &y)
	*y = 5
	fmt.Println("Value of x after", y)
	fmt.Println("Address of x after", &y)
}
