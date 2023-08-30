package main

import "testing"

func TestHello(t *testing.T ) {
	result := Hello("Chris")
	expected := "Salve, Chris"
	
	if result != expected {
		t.Errorf("got %q want %q", result, expected)
	}
}  