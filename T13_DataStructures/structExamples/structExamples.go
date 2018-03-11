package main

import (
	"fmt"
)

// Underlying type is a Struct.
type person struct {
	firstName string
	lastName  string
	age       int
}

// Inheritence (embedded types)
type student struct {
	person        // <----
	schoolClassID string
	var1          int
}

// Promotion
type csStudent struct {
	student
	var1 int //<----
	// Calling <mycsStudent>.var1 will use this instance of var1.
	// NOTE: We still store a separate variable also called var1
	// in the parent.
}

// Methods (Functions with Receivers, i.e. functions attached to types)
func (p person) getFullName() string {
	return (p.firstName + " " + p.lastName)
}

// Method Overrides (Another form of Promotion)
func (s student) getFullName() string {
	return ("Student Name Override")
}

// Tags:
// In this example, tags are used to specifiy
// that certain attributes should be renamed
// when being push to json strings.
type dataScore struct {
	ScoreVal int `json:"score"`
}

func main() {
	basicStructTests()
}

func basicStructTests() {
	s1 := student{person{"Mark", "Lorem", 20}, "1A", 3}
	s2 := student{person{"Jane", "Ipsum", 25}, "2A", 4}
	fmt.Println(s1.firstName, s1.lastName, s1.age)
	fmt.Println(s2.firstName, s2.lastName, s2.age)

	var s3 student
	fmt.Println("student 3: ", s3)

	var s4 = csStudent{
		student: student{
			person:        person{"Censu", "Farru", 27},
			schoolClassID: "3A",
			var1:          4},
		var1: 5}
	fmt.Println("s4.var1:'", s4.var1, "'")
	fmt.Println("s4.student.var1:'", s4.student.var1, "'")

	fmt.Printf("s4.getFullName():'%s'\n", s4.getFullName())
}
