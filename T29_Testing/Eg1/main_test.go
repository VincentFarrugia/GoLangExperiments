/*
An example of a basic GoLang UnitTest.
Test Functions:
- Should have one parameter: t *testing.T
- Should have names with the format "Test<NameOfFunctionUnderTest>"
- Should call t.Error or t.Fail to indicate a failure. (t.Error provides more detail.)
- Should call t.Log to provide non-failing information.
- Must be saved in a file named with the format "<something>_test.go"
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
