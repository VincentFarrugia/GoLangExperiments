package main

import (
	"fmt"
)

func main() {

	a := 43

	// REFERENCING OPERATOR, &
	fmt.Println("address of a - ", &a)
	fmt.Println("value of a - ", a)

	// DEREFERENCING OPERATOR, *
	ptrToA := &a
	fmt.Println("dereferencing ptr to a: ", *ptrToA)

	// POINTER TO AN INT
	var ptrToInt *int = &a
	fmt.Println(ptrToInt)

	zero_this_var_value(ptrToA)
	fmt.Println("Value of a after zero func: ", ptrToA)
}

func zero_this_var_value(x *int) {
	*x = 0
}
