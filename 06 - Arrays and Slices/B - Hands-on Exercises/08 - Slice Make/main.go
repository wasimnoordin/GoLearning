/*

For this exercise, do the following:
● Create a slice to store the names of all of the states in the United States of America.
○ Use make and append to do this.
○ Goal: do not have the array that underlies the slice created more than once.
● Print out
○ the len
○ the cap
○ the values, along with their index position, without using the range clause.
● Here is a list of the 50 states:
` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `
Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `
Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `
Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `
Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,

*/

package main

import (
	"fmt"
)

func main() {

	// when we create a slice using the length=capacity, any append will be first put into the nth index
	x1 := make([]int, 50) // len & cap = 50 -> will put zero values into slice
	fmt.Println(x1)

	// when we create a slice using 0 for the length, any append will be first put into the 0th index
	x2 := make([]int, 0, 50) //len = 0, cap = 50 -> will not put any values in
	fmt.Println(x2)
	fmt.Println("---------------------------------------")

	// now lets do the exercise

	xs := make([]string, 0, 50)
	xs = append(xs, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `	Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	fmt.Println(len(xs))
	fmt.Println(cap(xs))

	for i := 0; i < len(xs); i++ {
		fmt.Println(xs[i], i)
	}
}
