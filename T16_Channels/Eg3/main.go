package main

import (
	"fmt"
	"math/big"
)

func main() {
	inputNum := 4
	//f := factorial(inputNum)
	f := calcFactorialWithGoRoutines(inputNum)
	fmt.Printf("Factorial of %d is: %v\n", inputNum, f)
}

func calcFactorialWithGoRoutines(n int) *big.Float {

	retVal := big.NewFloat(1)
	numCalcsPerGoRoutine := 5
	divRes := float64(n) / float64(numCalcsPerGoRoutine)
	totalNumWorkersRequired := int(divRes)
	if (divRes - float64(totalNumWorkersRequired)) > 0 {
		totalNumWorkersRequired++
	}

	subTaskOutChan := make(chan *big.Float)
	rangeMin := 0
	limit := n
	maxSimultaneousWorkers := 8
	currentNumWorkers := 0
	bAllWorkDone := false
	for !bAllWorkDone {
		// Create new workers if necessary.
		if (totalNumWorkersRequired > 0) && (currentNumWorkers < maxSimultaneousWorkers) {
			numNewWorkersToCreate := minInt(totalNumWorkersRequired, maxSimultaneousWorkers)
			for i := 0; i < numNewWorkersToCreate; i++ {
				rangeMax := minInt(rangeMin+1+numCalcsPerGoRoutine, limit)
				go factorialSubWorker(rangeMin+1, rangeMax, subTaskOutChan)
				rangeMin = rangeMax
				currentNumWorkers++
				totalNumWorkersRequired--
			}
		}
		// Receive results.
		bAllWorkDone = ((currentNumWorkers == 0) && (rangeMin >= limit))
		if !bAllWorkDone {
			retVal = retVal.Mul(retVal, <-subTaskOutChan)
			currentNumWorkers--
		}
	}

	return retVal
}

func factorialSubWorker(rangeStart int, rangeEnd int, resultChan chan *big.Float) {
	result := big.NewFloat(1)
	counter := big.NewFloat(float64(rangeStart))
	for i := rangeStart; i <= rangeEnd; i++ {
		result = result.Mul(result, counter)
		counter.SetInt64(int64(i + 1))
	}
	resultChan <- result
}

func minInt(x int, y int) int {
	if x <= y {
		return x
	}
	return y
}

// Version without go routines.
func factorial(n int) int {
	retVal := 1
	for i := n; i > 0; i-- {
		retVal *= i
	}
	return retVal
}

func factorialSubWorkerNoChan(rangeStart int, rangeEnd int) int {
	result := 1
	for i := rangeStart; i <= rangeEnd; i++ {
		result *= i
	}
	return result
}
