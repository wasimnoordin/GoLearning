package main

import "fmt"

func main() {
	fmt.Println(printSquare(square, 5))

}

func printSquare(function func(int) int, x int) string {
	y := function(x)
	return fmt.Sprintf("the number %v squared is %v", x, y)
}

func square(n int) int {
	return n * n
}
