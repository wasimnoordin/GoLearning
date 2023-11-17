package main

import "fmt"

// dog struct represents a dog with a name.
type dog struct {
	first string
}

// walk is a method with value receiver; it prints the dog's action of walking.
func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walking.")
}

// run is a method with pointer receiver; it changes the dog's name and prints its action of running.
func (d *dog) run() {
	d.first = "Rover" // Change the dog's name to "Rover".
	fmt.Println("My name is", d.first, "and I'm running.")
}

func main() {
	// Creating a dog instance and calling its methods.
	d1 := dog{"Henry"} // Initialize a dog named "Henry".
	d1.walk()          // Henry walks.
	d1.run()           // Henry's name is changed to Rover, then he runs.

	// Creating a dog instance using a pointer and calling its methods.
	d2 := &dog{"Padget"} // Initialize a pointer to a dog named "Padget".
	d2.walk()            // Padget walks. Method call is automatically dereferenced.
	d2.run()             // Padget's name is changed to Rover, then he runs.
}
