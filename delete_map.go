package main

import (
	"fmt"
)

func main() {
	mymap := map[string]int{"Date of year": 1990, "born date": 14, "Born day": 5}
	fmt.Println(mymap)
	fmt.Printf("Deleting Born day from map ")
	delete(mymap, "Born day")
	fmt.Printf("New map list\n")
	fmt.Println(mymap)
}
