package main

import (
	"fmt"
	"sync"
	"time"
)

// Concurrency is achieved by go routines.
// In the example below we have a program which uses
// 3 go routines: main, foo and bar.
// We must tell main to wait for the other go routines
// to complete before exiting the main program.
// We do this using WaitGroups.

// WaitGroups are essentially semaphores.
// Counters for resource use.

var wg sync.WaitGroup

func main() {
	//noConcurrencyExample()
	concurrencyExample()
}

func noConcurrencyExample() {
	foo()
	bar()
}

func concurrencyExample() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i <= 100; i++ {
		fmt.Println("Foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i <= 100; i++ {
		fmt.Println("Bar:", i)
	}
	wg.Done()
}

func barWithPause() {
	for i := 0; i <= 100; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}
