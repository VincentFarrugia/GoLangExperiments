///////////////////////////////////////////////////////////////////
// Exercise description:
// Write a function with one variadic parameter that finds the greatest
// number in a list of numbers.
///////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

func main() {
	largestNum, _ := findLargetsNum(3, 4, 6, 1, 103, 15)
	fmt.Println("Largets number out of 3,4,6,1,103,15 is:", largestNum)
}

func findLargetsNum(a ...int) (int, bool) {

	if len(a) < 0 {
		// Empty list. Return error.
		return -1, true
	}

	largestInt := a[0]

	for i := 0; i < len(a); i++ {
		if a[i] > largestInt {
			largestInt = a[i]
		}
	}

	return largestInt, false
}
