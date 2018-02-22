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
}
