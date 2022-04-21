package main

import (
	"fmt"
)

var x int = 11

func main() {
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	shift := x << 1
	fmt.Printf("%d\t%b\t%#x\n", shift, shift, shift)
}
