package main

import (
	"fmt"
)

func main() {
	x := []int{1, 24, 45, 22}
	fmt.Println(x)
	x = append(x, 33, 44, 55)
	fmt.Println(x)
	y := []int{100, 1001, 10003}
	x = append(x, y...)
	fmt.Println(x)

}
