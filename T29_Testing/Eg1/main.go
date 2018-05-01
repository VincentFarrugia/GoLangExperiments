package main

import (
	"fmt"
)

func main() {
	numA := 10
	numB := 5
	fmt.Printf("%d + %d = %d\n", numA, numB, Sum(numA, numB))
}

// Sum takes two integers and returns the result of their addition.
func Sum(x int, y int) int {
	return (x + y)
}
