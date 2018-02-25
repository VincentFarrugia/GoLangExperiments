///////////////////////////////////////////////////////////////////
// Exercise description:
// Modify E8 to use a func expression.
///////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

func main() {

	inNum := 1
	fmt.Print("Enter a number: ")
	fmt.Scan(&inNum)

	// halfFunc is a function expression.
	halfFunc := func(x int) (halfVal float32, argWasEven bool) {
		halfVal = (float32(x) / 2.0)
		argWasEven = ((x % 2.0) == 0)
		return halfVal, argWasEven
	}

	halfOfInput, inputWasEven := halfFunc(inNum)
	fmt.Printf("Half of %v is %v\n", inNum, halfOfInput)
	if inputWasEven {
		fmt.Printf("Input %v was even\n", inNum)
	} else {
		fmt.Printf("Input %v was NOT even\n", inNum)
	}
}
