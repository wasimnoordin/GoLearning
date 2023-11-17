package main

import "fmt"

func main() {
	a := 1
	fmt.Println(&a) // pointer sematics, moved to heap
	// fmt.Println(a) // value sematics, moved to stack
}

/*
escapes to the heap
variable shared between main() and Println()

moved to heap
variable moved to the heap
*/
