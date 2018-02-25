package main

import (
	"fmt"
)

// MAIN is the entry point to your program.
func main() {

}

// Simple function with one parameter and no return type ("void").
func myFuncA(v string) {
	fmt.Println(v)
}

// Function which takes two parameters and no return type.
func myFuncB(v1 string, v2 int) {
	fmt.Println(v1, v2)
}

// Function which takes two strings,
// copies them to a string and returns the resulting string.
func myFuncC(v1, v2 string) string {
	return fmt.Sprint(v1, v2)
}

// Function with named return.
// (Not really necessary if we only have one return parameter)
func myFuncD(v1 string) (retStr string) {
	retStr = v1
	return
}

// Function with multiple return parameters.
func myFuncE(v1, v2 string) (string, string) {
	return (fmt.Sprint(v1, v2)), (fmt.Sprint(v2, v1))
}

// Variatic function.
// Go Docs:
// "The final parameter in a function signature may have a type prefixed with ...
// A function with such a parameter is called variadic and may be invoked with
// zero or more arguments for that parameter"
func myFuncF(sf ...float64) float64 {
	// Calculates the average of all the floats inputted.
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

// Variatic Arguments passed into a Variatic Function
func myFuncG() {
	// Here, data is one structure (1 slice).
	data := []float64{43, 56, 87, 10, 43, 58}
	// Here, we open up the data slice structure into individual items.
	myFuncF(data...)
}

// Anonymous function being assigned to a Function Expression
// and having the Function Expression being called.
func myFuncH() {
	funcExp := func() {
		fmt.Println("Hello world!")
	}
	funcExp()

	// Cool, you can do this.
	// Prints out "func()" as the type.
	fmt.Printf("%T\n", funcExp)
}

// Function which returns a function.
func myFuncI() func() string {
	return func() string {
		return "Hello world!"
	}
}

// Function which takes a function as a parameter.
// (Callbacks)
func myFuncJ(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

// Example of callback usage.
// A Filter function which gets passed the list of items (slice)
// and also the algorithm used to determine if an element should be kept or not.
func filter(numbers []int, filterTest func(int) bool) []int {
	xs := []int{}
	for _, n := range numbers {
		if filterTest(n) {
			xs = append(xs, n)
		}
	}
	return xs
}

// Recursive function.
func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

// Defer.
// Any statements with the defer keyword
// will run right before the parent function exits.
// Useful for closing files/streams or other heavy resources.
func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}

func deferExample() {
	defer world()
	hello()
}

// Every function in Go uses pass-by-value for its parameters.
// For reference types, a COPY of the reference is passed to the function.
// Eg. A slice's contents are NOT copied over but a reference to the slice is coppied/passed-by-value to the function.

func myFuncK() {
	m := make([]string, 1, 25)
	fmt.Println(m) // [ ]
	changeSlice(m)
	fmt.Println(m) // [Censu]
}

func changeSlice(z []string) {
	z[0] = "Censu"
	fmt.Println(z) // [Censu]
}

// Anonymous self executing function.
// (Function with no name and runs/executes itself)
func myFuncL() {
	func() {
		fmt.Println("I'm the anonymous self executing function!")
	}()
	// Adding the brackets at the end will call the anonymous function.
}
