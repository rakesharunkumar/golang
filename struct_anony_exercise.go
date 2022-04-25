package main

import (
	"fmt"
)

func main() {
	anonymous1 := struct {
		name      string
		age       int
		fav_color map[int]string
		fav_car   []string
	}{
		name:      "MS",
		age:       35,
		fav_color: map[int]string{1: "Green", 2: "Red"},
		fav_car:   []string{"Benz", "Jaguar", "Range rover"},
	}

	fmt.Println(anonymous1)
	for i, v := range anonymous1.fav_color {
		if i == 2 {
			fmt.Println(i, v)
		}
	}
	fmt.Printf("List of fav cars\n")
	for _, v := range anonymous1.fav_car {
		fmt.Println(v)
	}
}
