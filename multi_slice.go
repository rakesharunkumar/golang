package main

import (
	"fmt"
)

func main() {
	person1 := []string{"Raam", "A", "23", "Singapore"}
	person2 := []string{"Laxman", "B", "20", "Taiwan"}
	fmt.Println(person1)
	fmt.Println(person2)
	combine_person := [][]string{person1, person2}
	fmt.Println(combine_person)
	fmt.Println(combine_person[0])
	fmt.Println(combine_person[1])
}
