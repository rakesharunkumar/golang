package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello, world")
	fmt.Fprintln(os.Stdout, "Hello, world")   //from os package library
	io.WriteString(os.Stdout, "Hello, world") //from io package library
}
