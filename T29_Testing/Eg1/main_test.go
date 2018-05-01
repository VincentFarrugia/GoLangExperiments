/*
An example of a basic GoLang UnitTest.

[Test Functions]:
- Should have one parameter: t *testing.T
- Should have names with the format "Test<NameOfFunctionUnderTest>"
- Should call t.Error or t.Fail to indicate a failure. (t.Error provides more detail.)
- Should call t.Log to provide non-failing information.
- Must be saved in a file named with the format "<something>_test.go"

[Launching Tests]:
- Use "go test" in the top directory to run _test files in that directory.
- Use "go test -run ^TEST-FUNCTION-NAME$" to run a specific test function
- Use "go test github.com/<user>/<reponame>" to run tests using a fully-qualified package name.
- Use the -v option param to see verbose output for your tests.

[Generating an HTML coverage report]
Run these commands to visualise which parts of your program have been covered
by the test functions and which statements were not tested.
	"go test -cover -coverprofile=c.out"
	"go tool cover -html=c.out -o coverage.html"


*/

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSum is a UnitTest for the basic Sum function.
// Here we use t.Errorf and t.Logf to report results.
func TestSum(t *testing.T) {
	numA := 21
	numB := 4
	total := Sum(numA, numB)
	expectedTotal := 25
	if total != expectedTotal {
		t.Errorf("TestSum failed! %d + %d = %d is incorrect. Should have got result: %d.", numA, numB, total, expectedTotal)
	} else {
		t.Logf("TestSum success! %d + %d = %d.", numA, numB, total)
	}
}

// TestSumUsingAssertPkg is a UnitTest for the basic Sum function but using asserts.
func TestSumUsingAssertPkg(t *testing.T) {
	numA := 21
	numB := 4
	total := Sum(numA, numB)
	assert.Equal(t, 25, total, "TestSumUsingAssertPkg assert failed!")
}

// TestSumUsingTestTable is a UnitTest for the basic Sum function using a Test-Table of values.
func TestSumUsingTestTable(t *testing.T) {

	// A Test Table contains a series of entries.
	// Each entry is an individual test.
	// Each entry contains values for Input and Expected Output.

	testTable := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, testEntry := range testTable {
		total := Sum(testEntry.x, testEntry.y)
		if total != testEntry.n {
			t.Errorf("Sum of (%d+%d) was incorrect. Got: %d. Want: %d.", testEntry.x, testEntry.y, total, testEntry.n)
		}
	}
}
