package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// This line prints "Hello Gophers!" to the standard output (console) using fmt.Println.
	// It automatically appends a newline at the end.
	fmt.Println("Hello Gophers!")

	// This line also prints "Hello Gophers!" to the standard output,
	// but it uses fmt.Fprintln and explicitly specifies os.Stdout as the output destination.
	// os.Stdout is a file object that represents the standard output.
	// Like fmt.Println, fmt.Fprintln appends a newline at the end.
	fmt.Fprintln(os.Stdout, "Hello Gophers!")

	// This line prints "Hello Gophers!" to the standard output using io.WriteString.
	// Here, os.Stdout is again used as the writer. io.WriteString does not append a newline,
	// so the output will be exactly what's in the string.
	io.WriteString(os.Stdout, "Hello Gophers!")
}
