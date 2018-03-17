package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Parallelism example.

// Special function to do program initialisation
func init() {
	// Specify that the Go program can use all of the physical
	// CPUs and/or cores available.
	runtime.GOMAXPROCS(runtime.NumCPU())
}

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go doSomething()
	go doSomethingElse()
	wg.Wait()
}

func doSomething() {
	for i := 0; i <= 90; i++ {
		fmt.Println("Do something:", i)
	}
	wg.Done()
}

func doSomethingElse() {
	for i := 0; i <= 90; i++ {
		fmt.Println("Do something else:", i)
	}
	wg.Done()
}
