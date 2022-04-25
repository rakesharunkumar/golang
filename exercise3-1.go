package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 4, 5, 6}
	for i := range x {
		fmt.Println(i)
		fmt.Printf("%T", i)
	}
}
