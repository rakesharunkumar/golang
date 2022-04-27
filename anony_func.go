package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	//anonymous func.
	func() {
		fmt.Println("Anonymous print")
	}()

	//anonymous print with return
	func() string {
		fmt.Println("Return with value")
		return "Test"
	}()

}
