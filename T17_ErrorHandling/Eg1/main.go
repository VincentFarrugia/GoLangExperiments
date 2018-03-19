package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
)

func init() {
	// Redirect logs to the file "log.txt"
	// (By default the log package uses standard-error == terminal)
	logFD, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(logFD)
}

func main() {
	_, err := os.Open("myfile.txt")
	if err != nil {

		// Prints to standard-output
		//fmt.Println("Error:", err)

		// Writes to standard-error and prints date-time of each logged message
		//log.Println("Error:", err)

		// Writes to standard-error but then causes the program to immediately exit.
		//log.Fatalln("Error:", err)

		// Panic causes the program to exit but also displays some usefaul debug info
		// like a callstack.
		//panic(err)
		//log.Panicln("Error:", err)
	}
}

// Example of creating a custom error value.
func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("MyError: Cannot perform square root of a negative number")
	}
	return math.Sqrt(f), nil
}

// Something like this can be used for specific repeat errors instead of duplicating code.
// (Prefix with 'Err' instead of 'err' to remain ideomatic when exporting)
var (
	errMyCustomError  = errors.New("MyCustomError: problem occurred")
	errMyCustomErrorB = errors.New("MyCustomErrorB: problem B")
)

// SIDE NOTE:
// Use fmt.Errorf("MyError: Some value %d", relevantInt)
// to create an error value with more detailed context information.

// Example of creating a custom error type with a struct.
type pathError struct {
	Op   string
	Path string
	Err  error
}

func (pe *pathError) Error() string {
	return fmt.Sprintf("Path Error: Op:'%s', Path:'%s', BaseError:'%v'", pe.Op, pe.Path, pe.Err.Error())
}
