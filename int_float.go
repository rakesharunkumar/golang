package main

import (
	"fmt"
)

var x int = 12
var y float32 = 13.245

func main() {
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	fmt.Printf("%T\n", y)
	fmt.Println(y)
}
