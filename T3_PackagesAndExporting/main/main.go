package main

import (
	"fmt"

	"github.com/VincentFarrugia/GoLangExperiments/T3_PackagesAndExporting/stringutils"
)

func main() {
	testStr := "Hello"
	fmt.Printf("Original string is: '%s'\n", testStr)
	testStr = stringutils.Reverse(testStr)
	fmt.Printf("Reversed string is: '%s'\n", testStr)
}
