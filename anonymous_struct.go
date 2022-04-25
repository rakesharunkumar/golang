package main

import "fmt"

// No need to declare type and name for anonymous.
func main() {
	p1 := struct {
		fname string
		lname string
	}{
		fname: "Sachine",
		lname: "Tendulkar",
	}
	fmt.Println(p1)
}
