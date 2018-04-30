// UNIT TESTS FOR THE BOOK STRUCT

package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Name: "The C Programming Language", Author: "Brian Kernighan and Dennis Ritchie", ISBN: "0-13-110362-8"}
	jsonBytes := book.ToJSON()
	assert.Equal(t, `{"Name":"The C Programming Language","Author":"Brian Kernighan and Dennis Ritchie","ISBN":"0-13-110362-8"}`,
		string(jsonBytes), "Book JSON marshalling failed!")
}

func TestBookFromJSON(t *testing.T) {
	jsonStr := `{"Name":"The C Programming Language","Author":"Brian Kernighan and Dennis Ritchie","ISBN":"0-13-110362-8"}`
	jsonBytes := []byte(jsonStr)
	book := FromJSON(jsonBytes)
	assert.Equal(t, Book{Name: "The C Programming Language", Author: "Brian Kernighan and Dennis Ritchie", ISBN: "0-13-110362-8"},
		book, "Book JSON unmarshalling failed!")
}
