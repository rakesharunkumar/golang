package main

import (
	"fmt"
)

func main() {
	x := 1
	fmt.Println(x)
	fmt.Println(&x)
	// Now print the type of variables.
	fmt.Printf("%T\n", x)  // This will print int type
	fmt.Printf("%T\n", &x) // This will print Type of &x which is pointer type.

	// Now declare some variable with type and print

	y := &x
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	//Now try to assign new variable of type int which as address.
	//var z int = &x
	// fmt.Println(z)
	// fmt.Printf("%T", z) //This will throw an error since &x is not pointer type.
	// so to make it correct we need to assign z as pointer type as below.
	var z *int = &x
	fmt.Printf("%T\n", z)
	fmt.Println(*y) //Deferencing the address which is referring to value of that address.
}
