package main

import (
	"fmt"
	"math/big"
	"sync"
)

// Job struct for factorials
type factorialJob struct {
	inputNum int
	result   *big.Float
}

func main() {

	// FAN OUT
	c := gen()
	numWorkers := 4
	workerOutChanSlice := make([]<-chan *factorialJob, 0)
	for i := 0; i < numWorkers; i++ {
		workerOutChanSlice = append(workerOutChanSlice, factorial(c))
	}

	// FAN IN
	resultChan := fanInResult(workerOutChanSlice...)

	// Output the result channel.
	numResults := 0
	for processedJob := range resultChan {
		fmt.Printf("Factorial of %d is: %v\n", processedJob.inputNum, processedJob.result)
		numResults++
	}
	fmt.Printf("Found %d results\n", numResults)
}

// func gen
// Generate 1000 numbers to perform factorial on
// and place onto a channel.
func gen() <-chan *factorialJob {
	out := make(chan *factorialJob)
	go func() {
		for i := 0; i < 10000; i++ {
			for j := 5; j < 15; j++ {
				nwJob := new(factorialJob)
				nwJob.inputNum = j
				nwJob.result = big.NewFloat(-1.0)
				out <- nwJob
			}
		}
		close(out)
		fmt.Println("**** Generator exiting... ****")
	}()
	return out
}

// func factorial
// Worker function to grab numbers from a channel,
// compute the factorial and output to another channel.
func factorial(inputChan <-chan *factorialJob) chan *factorialJob {
	out := make(chan *factorialJob)
	go func() {
		for n := range inputChan {
			nwRes := fact(n.inputNum)
			n.result.Set(nwRes)
			out <- n
		}
		close(out)
		fmt.Println("**** Factorial worker exiting... ****")
	}()
	return out
}

// func fanInResult
// Performs a FAN-IN by merging all worker output channels
// into one result channel.
func fanInResult(workerOutChannels ...<-chan *factorialJob) chan *factorialJob {
	out := make(chan *factorialJob)
	var wg sync.WaitGroup
	wg.Add(len(workerOutChannels))
	go func() {
		for _, workerOutChan := range workerOutChannels {
			for x := range workerOutChan {
				out <- x
			}
			wg.Done()
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// func fact
// A function which computes the factorial of a given number.
func fact(x int) *big.Float {
	result := big.NewFloat(1)
	counter := big.NewFloat(float64(x))
	for i := x; i > 0; i-- {
		result = result.Mul(result, counter)
		counter = counter.SetFloat64(float64(i - 1))
	}
	return result
}
