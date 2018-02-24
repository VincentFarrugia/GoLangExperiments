///////////////////////////////////////////////////////////////////
// Exercise description:
// Create a program that prints to the terminal asking for a user to enter
// a small number and a larger number. Print the remainder of the larger
// number divided by the smaller number.
///////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

func main() {

	inNumA := 1
	inNumB := 1
	fmt.Print("Enter a number: ")
	fmt.Scan(&inNumA)
	fmt.Print("Enter a larger number: ")
	fmt.Scan(&inNumB)
	fmt.Printf("Remainder from %d/%d is: %v", inNumB, inNumA, (inNumB % inNumA))
}
