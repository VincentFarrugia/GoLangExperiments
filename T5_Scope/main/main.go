package main

// [FILE LEVEL SCOPE]
// fmt is imported just for this file.
import (
	"fmt"
)

// [SAME-PACKAGE LEVEL SCOPE]
// x can be used in both 'main', 'doSomething'
// and also in other files within the same package as this file.
var x int = 33

// [EXTERNAL-PACKAGE LEVEL SCOPE]
// Can be used within this package
// but also to other go files which import this package.
var W string = "test"

func main() {
	// [FUNCTION LEVEL SCOPE]
	y := 10
	fmt.Println(x)
	fmt.Println(y)
	doSomething()
}

func doSomething() {
	fmt.Println(x)

	// Invalid, since y is only accessible in main func scope above.
	//fmt.Println(y)

	// Invalid, since z is not accessible at this point.
	//fmt.Println(z)
	//z := "MyString"

	{
		// [BLOCK LEVEL SCOPE]
		B := "To be or not to be"
		fmt.Println(B)
	}

	// UNIVERSE LEVEL SCOPE
	// (similar to language keywords)
	// You can use the keyword "true", anywhere.
	flag := true
	fmt.Println(flag)

	// This is possible because P is at PACKAGE SCOPE.
	// even if it is not declared at the top of this file.
	fmt.Println(p)

	{
		// [VARIABLE SHADOWING]
		// (Not good practice)
		flag := false
		fmt.Println(flag)
	}

	// Anonymous function.
	// (Function without a name)
	// Function expression.
	// (Assigning a function to a variable)
	increment := func(val int) int {
		return (val + 1)
	}
	m := 34
	increment(m)

	// Function expression again.
	somethingFunc := wrapper()
	somethingFunc()
}

// Function that returns a function
func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

// PACKAGE SCOPE, even if this is not at the top of the file.
var p float32 = 3.4
