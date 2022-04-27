package main

import (
	"fmt"
)

func main() {
	defer func1()
	func2()

}
func func1() {
	fmt.Println("from func1")
}

func func2() {
	fmt.Println("from func2")
}
