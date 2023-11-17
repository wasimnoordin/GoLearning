package main

import "fmt"

func main() {
	x := sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(x)

}

func sum(ii []int) int {
	total := 0
	for _, v := range ii {
		total += v
	}
	return total
}

// Mame returns, not advised to use as they lack clarity:

/*

func sum(ii []int) (total int) {
	total = 0
	for _, v := range ii {
		total += v
	}
	return
}
*/
