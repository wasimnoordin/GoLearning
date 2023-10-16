package main

import "fmt"

type ByteSize int // to avoid the overflow, use uint64

const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // ByteSize is a new type we are creating (self-documenting)
	MB
	GB
	TB
	PB
	EB
)

func main() {
	fmt.Printf("%d \t \t \t %b \n", KB, KB)
	fmt.Printf("%d \t \t %b \n", MB, MB)
	fmt.Printf("%d \t \t %b \n", GB, GB)
	fmt.Printf("%d \t \t %b \n", TB, TB)
	fmt.Printf("%d \t %b \n", PB, PB)
	fmt.Printf("%d \t %b \n", EB, EB)
}
