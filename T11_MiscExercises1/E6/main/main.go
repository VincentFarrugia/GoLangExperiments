///////////////////////////////////////////////////////////////////
// Exercise description:
// Write a program that prints the numbers from 1 to 100. But for
// multiples of three print "Fizz" instead of the number and for the
// multiples of five print "Buzz". For numbers which are multiples of both
// three and five print "FizzBuzz".
///////////////////////////////////////////////////////////////////

package main

import "fmt"

func main() {

}

func versionA() {
	for i := 1; i <= 100; i++ {
		if ((i % 5) == 0) && ((i % 3) == 0) {
			fmt.Println("FizzBuzz")
		} else if (i % 3) == 0 {
			fmt.Println("Fizz")
		} else if (i % 5) == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func versionB() {

	str := ""
	for i := 1; i <= 100; i++ {

		str = ""
		if (i % 3) == 0 {
			fmt.Println("Fizz")
		}

		if (i % 5) == 0 {
			fmt.Println("Buzz")
		}

		if str == "" {
			fmt.Println(i)
		}
	}
}
