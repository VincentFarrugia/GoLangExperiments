///////////////////////////////////////////////////////////////////
// Exercise description:
// Print all of the even numbers between 0 and 100
///////////////////////////////////////////////////////////////////

package main

import "fmt"

func main() {

	fmt.Println("All even numbers between 0 and 100")
	for i := 0; i <= 100; i++ {
		if (i % 2) == 0 {
			fmt.Println(i)
		}
	}
}
