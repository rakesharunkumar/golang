package main

import (
	"fmt"
)

var i [10]int

func main() {
	fmt.Println(i)
	i[4] = 12
	fmt.Println(i[4])
}
