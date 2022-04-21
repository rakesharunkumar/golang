package main

import (
	"fmt"
)

func main() {
	s := "Test message"
	fmt.Println(s)
	fmt.Printf("%T", s)
	tobyte := []byte(s)
	fmt.Println(tobyte)
	t := len(s)
	fmt.Println(t)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
}
