package main

import "fmt"

func main() {
	x := []string{"sita", "gita", "neeta"}
	y := []string{"raam", "laxman", "ganesh"}

	fmt.Println(x)
	fmt.Println(y)

	z := [][]string{x, y}
	fmt.Println(z)

	for i, v := range z {
		fmt.Println(i, v)
		for j, k := range v {
			fmt.Println(j, k)
		}
	}

}
