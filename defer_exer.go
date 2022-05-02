package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	defer foo()
	bar()
}

func foo() {
	fmt.Println("From foo")
}

func bar() {
	fmt.Println("From bar")
}
