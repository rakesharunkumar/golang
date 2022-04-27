package main

import (
	"fmt"
)

func main() {
	printname()
	// pass the value
	sendvalue("Rahul")
	// call by variable value, needs return string in func.
	return_func := stringfunc("Dravid")
	fmt.Println(return_func)

	//Multiple value
	a, b := multi_return("Sachin", 40)
	fmt.Println(a, b)
}

func printname() {
	fmt.Println("Printing from Name")
}

func sendvalue(s string) {
	fmt.Printf("My Name is %s", s)
}
func stringfunc(sentstring string) string {
	fmt.Println("From return function")
	return (sentstring)
}
func multi_return(fn string, age int) (string, int) {
	fmt.Printf("His firstname is %s\n", fn)
	fmt.Printf("and he is %d year old", age)
	return fn, age
}
