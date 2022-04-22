package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 6, 34}
	fmt.Println(x)
	x = append(x[:2], x[:3]...)
	fmt.Println(x)
}
