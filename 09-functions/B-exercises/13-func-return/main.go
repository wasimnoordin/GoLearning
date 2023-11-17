package main

import "fmt"

func main() {
	fn := outerFunc()
	fmt.Println(fn())

}

func outerFunc() func() int {
	return func() int {
		return 42
	}
}
