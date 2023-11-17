package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123` // Original password

	// Generating a hashed version of the password using bcrypt.
	// bcrypt.MinCost is the cost of processing the password. Higher cost means more secure hash.
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)  // original password
	fmt.Println(bs) // hashed password

	loginPassword1 := `password1234` // Password attempt for login

	// Comparing the hashed password with the login attempt.
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPassword1))
	if err != nil {
		fmt.Println("YOU CAN'T LOGIN")
		return
	}

	fmt.Println("You're logged in")
}
