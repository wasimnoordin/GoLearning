package main

import "fmt"

func main() {
	summation()

}

var summation = func() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for _, num := range nums {
		sum += num
		fmt.Println(sum)
	}
}
