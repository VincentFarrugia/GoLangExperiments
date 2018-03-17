package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// MUTEX to prevent race conditions.

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go incrementWorker("WA:")
	go incrementWorker("WB:")
	wg.Wait()
	fmt.Println("Final Result:", counter)
}

func incrementWorker(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)

		// Curly brackets not necessary here.
		// They are just to make the critical section clearer.
		{
			mutex.Lock()
			x := counter
			x++
			counter = x
			fmt.Println(s, i, "Counter:", counter)
			mutex.Unlock()
		}
	}
	wg.Done()
}
