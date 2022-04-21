package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d\n", i)
	}
	j := 1
	for j < 3 {
		fmt.Printf("%U\n", j)
		j++
	}

}
