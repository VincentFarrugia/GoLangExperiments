package main

import "fmt"

func main() {

	// English unicode characters take 1 byte per letter.
	word1 := "Hello"
	fmt.Printf("Bytes for '%s': %v\n", word1, []byte(word1))

	// For something more complex, like Chinese characters,
	// up to 4 bytes can be used per character.
	// (UTF8 uses 1 to 4 bytes per character. i.e max of 32 bits)
	// This example uses 3 bytes per character.
	word2 := "世界"
	fmt.Printf("Bytes for '%s': %v\n", word2, []byte(word2))
	fmt.Println()

	// NOTE: []byte(string(i))
	// Converts integer to string.
	// Converts string into a slice of bytes.
	for i := 50; i <= 140; i++ {
		fmt.Printf("%d   -   %s   -   %v\n", i, string(i), []byte(string(i)))
	}

	fmt.Println()

	// NOTE: The first parameter of the Println
	// will not cause the letter 'i' to be printed.
	// this is because single quoted character denote
	// something of type Rune.
	//
	// Rune is an alias for int32 and is used for refering to Unicode characters.
	//
	j := 33
	fmt.Println('i', "-", string(j), "-", []byte(string(j)))
	fmt.Println()

	// %T will print out 'int32'
	// Rune is an alias for int32.
	foo := 'a' // THIS IS A RUNE
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
	fmt.Printf("%T \n", rune(foo))
	fmt.Println()

	foobar := "a"   // THIS IS A STRING
	foobarV2 := `a` // THIS IS A STRING
	fmt.Printf("FooBar: '%v'\n", foobar)
	fmt.Printf("FooBar_v2: '%v'\n", foobarV2)

	// A String using `` will NOT interpret escaped characters.
	// A String using "" will interpret escaped characters.

	// A string is effectively a collection of Runes.
}
