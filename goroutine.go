package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Test")
	wg.Add(2)
	go goroutine1()

	go goroutine2()
	wg.Wait()
	fmt.Println("Exit from pgm")

}

func goroutine1() {
	fmt.Println("From NUMCPU", runtime.NumCPU())
	fmt.Println("From Goroute1", runtime.NumGoroutine())
	wg.Done()
}

func goroutine2() {
	fmt.Println("From NUMCPU", runtime.NumCPU())
	fmt.Println("From Goroute2", runtime.NumGoroutine())
	wg.Done()
}
