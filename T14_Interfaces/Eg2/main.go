package main

import (
	"fmt"
	"sort"
)

// Common Interfaces:
// - io.Reader, io.Writer
// - http.Handler
// - sort.Interface

type people []string

func main() {

	// Using helper function for sorting slices of string.
	peopleNames := people{"Mark", "Andrew", "Peter", "Ian"}
	sort.Strings(peopleNames)
	fmt.Println("Sorted Names 1:", peopleNames)

	// Using generic sorting function for slices.
	peopleNames = people{"Tom", "Robert", "Phil", "John"}
	sort.Slice(peopleNames, func(i, j int) bool {
		return peopleNames[i] < peopleNames[j]
	})
	fmt.Println("Sorted Names 2:", peopleNames)

	// Using sorting function for any type with the sort.Interface interface.
	peopleNames = people{"Vincent", "Aaron", "Paul", "Mike"}
	sort.Sort(peopleNames)
	fmt.Println("Sorted Names 3:", peopleNames)

	// Sorting with descending order.
	peopleNames = people{"Noel", "James", "Bob", "Simon"}
	sort.Slice(peopleNames, func(i, j int) bool {
		return peopleNames[i] > peopleNames[j]
	})
	fmt.Println("Sorted Names 4 (Desc):", peopleNames)

	// Sorting a slice of int.
	myInts := []int{30, 2, 33, 417, 947296, 33, 3, 45, 67, 3, 5}
	sort.Ints(myInts)
	fmt.Println("Sorted Ints:", myInts)

	// Descending order of ints.
	sort.Sort(sort.Reverse(sort.IntSlice(myInts)))
	fmt.Println("Sorted Ints (Desc):", myInts)

	// Testing out "ToString" like functionality.
	var d = myCustomType{3, "Hello"}
	fmt.Printf("d to string:'%s'\n", d)
}

// Ascending Order Sort for custom type
// which implements the sort.Interface interface.
func (s people) Len() int           { return len(s) }
func (s people) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s people) Less(i, j int) bool { return s[i] < s[j] }

// Example of type implementing an interface
// similar to "ToString" in other languages.

type myCustomType struct {
	numA int
	strA string
}

func (d myCustomType) String() string {
	return fmt.Sprintf("%d%s", d.numA, d.strA)
	//return strconv.Itoa(d.numA) + string(d.strA)
}

// Empty Interface.
// The interface with no methods.
// Therefore, every type implements the Empty Interface.

func doWhatever(k interface{}) {

}

func doWhateverVariadic(k ...interface{}) {

}

// Assertion.
// Can be used to check if an interface is of type T.
// If successful, you can cast down to a more concrete type.

func typeAssertionExample() {
	var name interface{} = "Tom"
	str, ok := name.(string)
	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Println("Value was not a string.")
	}
}
