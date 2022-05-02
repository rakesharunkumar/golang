package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7}
	y := []int{11, 12, 13, 14, 15, 16, 17}
	fmt.Println(x)
	res := foo(x...)
	fmt.Println(res)
	//call func bar with slice of int.
	bar(y)
}

func foo(i ...int) int {
	total := 0
	for _, v := range i {
		total = total + v
	}
	return total
}

func bar(y []int) {
	fmt.Println(y[0])
	sum := 0
	for i := 0; i < len(y); i++ {
		fmt.Printf("Sum of %d + %d is \t", sum, y[i])
		sum = sum + y[i]
		fmt.Println(sum)
	}
	fmt.Println(sum)
}
