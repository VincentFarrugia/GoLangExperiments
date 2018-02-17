package main

import "fmt"

func main() {

	// Plain Decimal.
	fmt.Println(42)

	// Decimal, binary, octal, hex.
	fmt.Printf("%d - %b - %o - %x \n", 42, 42, 42, 42)

	// Basic loop displaying decimal, binary, hex for range.
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}

	// Loop which also displays utf-8 char.
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
