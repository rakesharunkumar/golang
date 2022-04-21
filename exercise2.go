package main

import (
	"fmt"
)

var x int = 42
var y string = "James bond"
var z bool = true

func main() {
	fmt.Println(x, y, z)
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
