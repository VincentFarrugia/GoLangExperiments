package main

import (
	"fmt"
)

func main() {

	//////////////////////////////////////////////////
	// Array (fixed contiguous memory)
	//////////////////////////////////////////////////
	var x [58]string
	fmt.Println(x)
	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}
	fmt.Println(x)

	//////////////////////////////////////////////////
	// Slice (Maps onto a range of elements of an underlying array)
	// (Similar to ArrayList...not LinkedList!)
	//////////////////////////////////////////////////

	// Making a slice from an array.
	var y []string = x[4:7]
	fmt.Println(y)

	// Create a Slice: SHORTHAND METHOD
	mySlice := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4]) // Slicing a slice.
	fmt.Println(mySlice[2])   // Access element by Index.

	// Create a Slice: VAR METHOD
	// (Remember, using var will initialise the variable with the default value.
	// In this case the default value for Slice, which is 'nil')
	var varSlice []int
	fmt.Println("VarSlice:", varSlice)
	fmt.Println("VarSlice is 'nil'?", (varSlice == nil))

	// Create a Slice: MAKE METHOD
	// Use 'make' to create a new slice initialised with a length and capacity.
	// The length will indicate the size of the slice (the window over the underlying array).
	// The capacity will indicate the default size of the underlying array.
	// Capacity will double every time length exceeds capacity. I.e whenever the underlying array becomes full.
	madeSlice := make([]int, 5, 10)
	madeSlice = append(madeSlice, 101)
	madeSlice = append(madeSlice, 200)
	fmt.Println("MadeSlice:", madeSlice)
	fmt.Println("Length of madeSlice is:", len(madeSlice))
	fmt.Println("Capacity of madeSlice is:", cap(madeSlice))

	// Access element by Index.
	// A string is made up of runes.
	// A rune is a unicode code point.
	// A unicode code point is 1-4 bytes.
	// Effectively, a string is a slice of bytes.
	// This prints out '83', which is unicode for 'S'
	fmt.Println("myString"[2])

	// If we use this instead of make,
	// the created slice will have
	// length and capacity equal to
	// the number of initialising elements
	testSlice := []int{101, 304, 425}
	fmt.Println("TestSlice:", testSlice)

	// Slicing a slice with a range. ([startIndex:endIndex]) End index is exclusive.
	// (Also note the extra comma at the end of the slice
	// if we want to write the initialiser list over multiple lines)
	namesSlice := []string{
		"Peter",
		"Paul",
		"Tom",
		"Mark",
		"John",
		"Ben",
	}
	fmt.Println("[1:2]", namesSlice[1:2])
	// Start from beginning of namesSlice.
	fmt.Println("[:2]", namesSlice[:2])
	// Range will end at the end of namesSlice.
	fmt.Println("[5:]", namesSlice[5:])
	// Take the entire range of namesSlice.
	fmt.Println("[:]", namesSlice[:])

	// Appending a Slice to a Slice.
	// (Notice the use of Variadic Argument)
	mergedSlice := append(madeSlice, testSlice...)
	fmt.Println("MergedSlice:", mergedSlice)

	// Delete using slice.
	// This effectively "deletes" slice entry index 2.
	weekdaysA := []string{"Monday", "Tuesday"}
	weekdaysB := []string{"Wednesday", "Thursday", "Friday"}
	theWeek := []string{}
	theWeek = append(theWeek, weekdaysA...)
	theWeek = append(theWeek, weekdaysB...)
	// The actual delete part.
	theWeek = append(theWeek[:2], theWeek[3:]...)

	// Multidimentional Slices:
	sliceOfSlicesOfString := make([][]string, 0)
	sliceOfSlicesOfString = append(sliceOfSlicesOfString, []string{"Hello", "world"})
	sliceOfSlicesOfString = append(sliceOfSlicesOfString, []string{"cool"})
	fmt.Println(sliceOfSlicesOfString)

	// Incrementing a Slice Item:
	// This is the same as: testSlice[0] += 1
	testSlice[0]++

	//////////////////////////////////////////////////
	// Map
	//////////////////////////////////////////////////

	//////////////////////////////////////////////////
	// Struct
	//////////////////////////////////////////////////

}
