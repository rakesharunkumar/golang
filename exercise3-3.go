package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x = append(x, 11)
	fmt.Println(x)
	x = append(x, 12, 13, 14, 15)
	fmt.Println(x)
	y := []int{16, 17, 18, 19, 20}
	z := append(x, y...)
	fmt.Println(z)
}
