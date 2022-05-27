package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup //use sync package wait group

func main() {
	fmt.Println("OS:\t", runtime.GOOS)
	fmt.Println("ARCH:\t", runtime.GOARCH)
	fmt.Println("No of CPU:\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
	//Now to call concurrancy func, we need to add go as prefix to make use of another cpu core.
	wg.Add(1)
	go concurence()
	fmt.Println("Go routine:\t", runtime.NumGoroutine()) //shows increament in goroutine
	wg.Wait()
}

func concurence() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	wg.Done()
}
