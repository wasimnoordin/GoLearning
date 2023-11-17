/*
○ create a func with the identifier foo that returns an int
○ create a func with the identifier bar that returns an int and a string
○ call both funcs
○ print out their results
*/

package main

import "fmt"

func foo() int {
	return 42
}

func bar() (int, string) {
	return 43, "bar says wassup"
}

func main() {
	fooInt := foo()
	barInt, barString := bar()

	fmt.Println("foo result:", fooInt)
	fmt.Println("bar result:", barInt, "\nbar string:", barString)
}
