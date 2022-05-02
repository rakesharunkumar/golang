package main

import (
	"fmt"
)

func main() {
	val := sum()
	fmt.Println(val())
}

func sum() func() int {
	return func() int {
		return 1 + 2
	}
}
