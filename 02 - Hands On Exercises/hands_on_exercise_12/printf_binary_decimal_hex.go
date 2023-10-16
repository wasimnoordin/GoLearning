/*

This program prints verbs to show the following numbers:
	- 747
	- 911
	- 90210

As:
	- demical
	- binary
	- hexadecimal

*/

package main

import "fmt"

func main() {
	x, y, z := 747, 911, 90210

	fmt.Printf("%d \t\t %b \t\t %#X\n", x, x, x) //decimal, binary, hex
	fmt.Printf("%d \t\t %b \t\t %#X\n", y, y, y) //decimal, binary, hex
	fmt.Printf("%d \t\t %b \t %#X\n", z, z, z)   //decimal, binary, hex
}
