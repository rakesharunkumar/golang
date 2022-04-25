package main

import (
	"fmt"
)

func main() {
	states := make([]string, 50, 50)
	fmt.Println(len(states))
	fmt.Println(cap(states))
	states = []string{`Karnataka`, `Andhra Pradesh`, `Tamil Nadu`, `Gujarath`}
	for v := range states {
		fmt.Println(states[v])
	}
}
