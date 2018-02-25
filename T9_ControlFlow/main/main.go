package main

import (
	"fmt"
)

func main() {

	// Basic For Loop
	for i := 1; i < 101; i++ {
		fmt.Println(i)
	}

	// Basic "While" Loop
	i := 0
	bIsDone := false
	for !bIsDone {
		i++
		fmt.Println(i)
		bIsDone = (i >= 5)
	}

	// Infinite Loop
	//for {
	//
	//}

	// Nested Loop
	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, " - ", j)
		}
	}

	// "For-each" Loop
	numArr := [5]int{0, 1, 2, 3, 4}
	sum := 0
	for idx, element := range numArr {
		sum += element
		fmt.Println("loop array element idx:", idx)
	}

	// "For-each" Loop where we do not care about the index.
	sum = 0
	for _, element := range numArr {
		sum += element
	}

	// Switch Statement
	// No default fallthrough
	// Fallthrough is optional and usage of 'break' is not needed.
	// Switch statements in Go also work with more than just integers.
	switch "NameB" {
	case "NameA":
		fmt.Println("Case with NameA")
	case "NameB":
		fmt.Println("Case with NameB")
		fallthrough // This will fall through to NameC but will NOT fall through to NameD after that.
	case "NameC":
		fmt.Println("Case with NameC")
	case "NameD":
		fmt.Println("Case with NameD")
	default:
		fmt.Println("Default case")
	}

	// Switch with multiple cases:
	switch "NameB" {
	case "NameA", "NameB":
		fmt.Println("Case with NameA or NameB")
	case "NameC", "NameD":
		fmt.Println("Case with NameC or NameD")
	default:
		fmt.Println("Default case")
	}

	// Switch with no expression:
	// Runs the first case which is true,
	// If none are true, then default is run.
	foobar := "Hello"
	switch {
	case len(foobar) == 2:
		fmt.Println("Foobar length is 2")
	case foobar == "World":
		fmt.Println("Foobar is 'World'")
	default:
		fmt.Println("Default case")
	}

	// Switching on type:
	/*x interface{}
	switch x.(type) {
	case int:
		fmt.Println("int")
	case float32:
		fmt.Println("float32")
	default:
		fmt.Println("x is of 'unknown type'")
	}*/
}
