package main

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	obtained := add(3, 4)
	expected := 7
	if obtained != expected {
		log.Fatalf("Error -  expected %v and obtained %v", expected, obtained)
	}
}

func TestSubtract(t *testing.T) {
	obtained := subtract(3, 4)
	expected := -1
	if obtained != expected {
		log.Fatalf("Error -  expected %v and obtained %v", expected, obtained)
	}
}
