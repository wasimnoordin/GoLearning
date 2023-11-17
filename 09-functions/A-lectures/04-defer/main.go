package main

import "fmt"

func main() {
	// Use the 'defer' keyword before calling the 'foo' function.
	// This means 'foo' will be deferred until the surrounding function (in this case, 'main') completes.
	// 'foo' will be executed just before 'main' exits.

	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
