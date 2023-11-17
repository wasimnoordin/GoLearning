/*
func syntax

a. no params, no returns
b. 1 param, no returns
c. 1 param, 1 return
d. 2 params, 2 returns

*/

// func (r reciever) identifier(p parameter(s)) (return(s)) { code }

package main

import "fmt"

func main() {
	foo()

	bar("Sim")

	s := chicanery("Heisenberg")
	fmt.Println(s)

	s1, n := mudamuda("Dio", 3)
	fmt.Println(s1, n)

}

// a.
func foo() {
	fmt.Println("I am from foo")
}

// b.
func bar(s string) {
	fmt.Println("My name is", s)
}

// c.
func chicanery(s string) string {
	return fmt.Sprint("Say my name... ", s)
}

// d.
func mudamuda(name string, punchesCombo int) (string, int) {
	punchesCombo *= 10
	return fmt.Sprint(name, " has a punch combo of"), punchesCombo
}
