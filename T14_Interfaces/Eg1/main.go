package main

import (
	"fmt"
	"math"
)

////////////////////
// SHAPE
////////////////////

// Shape A generic interface for all Shapes.
type Shape interface {
	perimeter() float64
	surfaceArea() float64
	volume() float64
}

////////////////////
// SQUARE
////////////////////

// Square A sub-type of Shape representing a Square
type Square struct {
	side float64
}

func (s Square) perimeter() float64 {
	return s.side * 4
}

func (s Square) surfaceArea() float64 {
	return s.side * s.side
}

func (s Square) volume() float64 {
	return 0
}

////////////////////
// CIRCLE
////////////////////

// Circle A sub-type of Shape representing a Circle
type Circle struct {
	radius float64
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) surfaceArea() float64 {
	return math.Pi * (math.Pow(c.radius, 2))
}

func (c Circle) volume() float64 {
	return 0
}

////////////////////

func printShapeInfo(s Shape) {
	fmt.Println(s)
	fmt.Println("Perimeter:", s.perimeter())
	fmt.Println("Surface Area:", s.surfaceArea())
	fmt.Println("Volumne:", s.volume())
}

func main() {
	s := Square{10}
	c := Circle{4.0}

	fmt.Println("Square S Info:")
	printShapeInfo(s)
	fmt.Println("Circle C Info:")
	printShapeInfo(c)
}

////////////////////
