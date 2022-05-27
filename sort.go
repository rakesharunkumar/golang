package main

import (
	"fmt"
	"sort"
)

func main() {
	intslice := []int{5, 3, 6, 1, 3, 66}
	stringslice := []string{"Sachin", "Sehwag", "Rahul", "Jadeja"}
	fmt.Println(intslice)
	sort.Ints(intslice)
	fmt.Println(intslice)

	fmt.Println(stringslice)
	sort.Strings(stringslice)
	fmt.Println(stringslice)
}
