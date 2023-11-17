package main

import (
	"log"
	"testing"
)

func TestParadise_vacation(t *testing.T) {
	obtained := paradise_vacation("Kenya")
	expected := "My idea of paradise is Kenya"
	if obtained != expected {
		log.Fatalf("error - obtained %v and expected %v", obtained, expected)
	}
}
