package main

import (
	"fmt"
	"math"
)

// circle represents a geometric circle with a radius.
type circle struct {
	radius float64
}

// shape is an interface that defines a method for calculating area.
type shape interface {
	area() float64
}

// area calculates the area of the circle.
// It implements the area method for the circle struct.
func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// info function takes a shape interface and prints its area.
func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	// Create a circle with a radius of 5.
	c := circle{5}

	// Print the area of the circle directly (without using the interface).
	fmt.Println(c.area())

	// The line below is commented out because it will not work due to a mismatch:
	// The method area() is defined on a pointer receiver for circle,
	// but when uncommented, it's invoked on a non-pointer value.

	// info(c)
}
