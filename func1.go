package main

import "fmt"

func main() {
	fmt.Println("This is the main function")
	simple_text()
	fmt.Println("End of main program")
}

func simple_text() {
	fmt.Println("This Message is from simple text function")
	for i := 0; i < 9; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
		//fmt.Println(i)
	}
}
