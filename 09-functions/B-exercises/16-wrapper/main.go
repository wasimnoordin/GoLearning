package main

import (
	"fmt"
	"time"
)

func main() {
	timeFunc(timeRange)

}

func timeRange() {
	for i := 0; i < 10000; i++ {
		fmt.Println(i)
	}
}

func timeFunc(f func()) { // callback as we pass a funcrtion as an argument
	start := time.Now()
	f() // wrapping the function before and after some processing
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
