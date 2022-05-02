package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	f1 := foo()
	fmt.Println(f1)
	f2, f3 := bar()
	fmt.Println(f2, f3)
}

func foo() int {
	return 3
}

func bar() (int, string) {
	return 32, "somevalue"
}
