package main

import (
	"fmt"
)

var x = []int{4, 5, 6, 7, 10}

func main() {
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Println("Now printing slice within slice")
	fmt.Println(x[1:]) // prints values from x[1] i,e from 5
	fmt.Println("Now printing value from  2nd to 4th position")
	fmt.Println(x[1:4])
	//printing index and value without using range.
	for j := 0; j < 4; j++ {
		fmt.Println(j, x[j])
	}
}
