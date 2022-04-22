package main

import (
	"fmt"
)

func main() {
	x := make([]int, 10, 20)
	fmt.Println(x)
	fmt.Println(cap(x))
	fmt.Println(len(x))
	x = append(x, 11)
	fmt.Println(x)
	fmt.Println(len(x))
}
