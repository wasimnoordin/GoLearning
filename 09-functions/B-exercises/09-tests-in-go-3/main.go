package main

import "fmt"

func main() {
	fmt.Println(paradise_vacation("Mauritius"))
}

func paradise_vacation(location string) string {
	return fmt.Sprint("My idea of paradise is ", location)
}
