package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	i := []int{1, 2, 3, 4, 5, 6}
	add := sum(i...)
	fmt.Println(add)
	//Pass sum function as arguemtn to another function to be called as callback.
	j := []int{7, 8, 9, 10}
	callback := even(sum, j...)
	fmt.Println(callback)
}

func sum(i ...int) int {
	total := 0
	for _, v := range i {
		//fmt.Println(v)
		total = total + v
	}
	return total
}

// Now pass func as a arguement which is called call back.

func even(f func(i ...int) int, j ...int) int {
	fmt.Println(j)
	fmt.Printf("%T\n", f)
	return f(j...) //returning the values of j to called function(callback)
}
