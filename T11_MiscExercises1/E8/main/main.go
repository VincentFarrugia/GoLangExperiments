///////////////////////////////////////////////////////////////////
// Exercise description:
// Write a function which takes an integer. The function will have two
// returns. The first return should be the argument divided by 2. The
// second return should be a bool that let’s us know whether or not the
// argument was even. For example:
// a. half(1) returns (0, false)
// b. half(2) returns (1, true)
///////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

func main() {

	inNum := 1
	fmt.Print("Enter a number: ")
	fmt.Scan(&inNum)
	halfOfInput, inputWasEven := half(inNum)
	fmt.Printf("Half of %v is %v\n", inNum, halfOfInput)
	if inputWasEven {
		fmt.Printf("Input %v was even\n", inNum)
	} else {
		fmt.Printf("Input %v was NOT even\n", inNum)
	}
}

func half(x int) (halfVal float32, argWasEven bool) {
	halfVal = (float32(x) / 2.0)
	argWasEven = ((x % 2.0) == 0)
	return halfVal, argWasEven
}
