package main

import (
	"fmt"
)

var x = []int{7, 8, 9, 10}
var z int = 10

func main() {
	//x := []int{1, 2, 4, 5}
	fmt.Println(x)
	fmt.Println(z)
	for i, v := range x {
		fmt.Println(i, v)
	}
}
