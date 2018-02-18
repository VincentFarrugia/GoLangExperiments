package main

import (
	"fmt"
)

// 1. Declaring a variable with a value,
// 2. Having it reusable within functions of current file/package
// 3. Having it exportable (starts with Capital letter)
var (
	D int = 23
)

func main() {

	// Shorthand for variable initialisation
	// Shorthand can only be used within functions.
	a := "Hello"
	b := 3
	fmt.Printf("Variable 'a' is of type '%T' and has value of '%v'\n", a, a)
	fmt.Printf("Variable 'b' is of type '%T' and has value of '%v'\n", b, b)

	// Declare variable with default value (aka zero-value)
	// This type of declaration can be used outside functions.
	var c int
	fmt.Printf("Default value of 'c' is '%v'\n", c)
	fmt.Printf("Default value of 'D' is '%v'\n", D)

	// Multiple initialisation. (Declaration + assignment)
	var e, f string = "str1", "str2"
	fmt.Printf("String e '%v', string f '%v'\n", e, f)
}

func DoSomething() {
	fmt.Println("Did something.")
}
