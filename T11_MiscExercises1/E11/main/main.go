///////////////////////////////////////////////////////////////////
// Exercise description:
// What's the value of this expression: (true && false) || (false && true) || !(false && false)?
///////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

func main() {
	// Result is TRUE
	bResult := ((true && false) || (false && true) || !(false && false))
	fmt.Printf("The result of '(true && false) || (false && true) || !(false && false) is: %v", bResult)
}
