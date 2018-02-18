package main

import "fmt"

// NOTE: See this for more interesting info.
// https://blog.golang.org/constants

// [EXAMPLE OF CONSTANT]
const p string = "test string"

// NEVER USE ALL CAPS FOR CONSTANT IDENTIFIERS
// (Unlike other programming languages)
// (Note: First letter can be caps if you wan to export the constant)

// [MULTIPLE INITIALIZATION]
const (
	Pi       = 3.14
	Language = "GO"
)

// [IOTA]
// Similar to Enums?
const (
	A = iota // 0
	B = iota // 1
	C = iota // 2
)

const (
	D = iota // 0
	E        // 1
	F        // 2
)

const (
	G = (10 + iota) // 10 + 0
	H               // 11 (10 + 1)
	I               // 12 (10 + 2)
)

const (
	Http_Ok                 = 200
	Http_Resource_Not_Found = 404
	Http_Server_Error       = 500
)

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

func main() {

	const q = 41
	fmt.Println("p - ", p)
	fmt.Println("q - ", q)

	fmt.Println("G - ", G)
	fmt.Println("H - ", H)
	fmt.Println("I - ", I)
}
