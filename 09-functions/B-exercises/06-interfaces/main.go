/*

● create a type SQUARE
○ length float64
○ width float64
● create a type CIRCLE
○ radius float64
● attach a method to each that calculates AREA and returns it
○ circle area= π r 2
■ math.Pi
■ math.Pow
○ square area = L * W
● create a type SHAPE that defines an interface as anything that has the AREA method
● create a func INFO which takes type shape and then prints the area
● create a value of type square
● create a value of type circle
● use func info to print the area of square
● use func info to print the area of circle

*/

package main

import (
	"fmt"
	"math"
)

// define SQUARE type
type SQUARE struct {
	length float64
	width  float64
}

// define CIRCLE type
type CIRCLE struct {
	radius float64
}

// attach area method to SQUARE
func (s SQUARE) area() float64 {
	return s.length * s.width
}

// attached area method to CIRCLE
func (c CIRCLE) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// define SHAPE interface
type SHAPE interface {
	area() float64
}

// define INFO function
func info(s SHAPE) {
	fmt.Printf("The area is: %.2f \n", s.area())
}

func main() {
	// create value of type SQUARE
	sq := SQUARE{
		length: 10,
		width:  10,
	}

	// create value of type CIRCLE
	ci := CIRCLE{
		radius: 5,
	}

	info(sq)
	info(ci)
}
