package main

import (
	"fmt"
)

func main() {
	a := 2
	b := 4
	c := 10
	sum := func1(a, b)
	fmt.Println(sum)
	//sending function as arguement to another function (callback function)
	res := func2(func1, c)
	fmt.Println(res)
	// fmt.Println()
}

func func1(i int, j int) int {
	return i + j
}
func func2(f1 func(i int, j int) int, c int) int {
	//fmt.Printf("%T\t", f1)
	x := f1(1, 5)
	return x + c
}
