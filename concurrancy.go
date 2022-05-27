package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS:\t", runtime.GOOS)
	fmt.Println("ARCH:\t", runtime.GOARCH)
	fmt.Println("Runtime:\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
	//Now to call concurrancy func, we need to add go as prefix to make use of another cpu core.

	go concurrancy()
	fmt.Println("Go routine:\t", runtime.NumGoroutine()) //shows increament in goroutine
}

func concurrancy() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}
