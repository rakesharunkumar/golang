package main

import (
	"fmt"
)

var x = 10
var y = `"Test
 message"`

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
