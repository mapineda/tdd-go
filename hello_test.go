package main

import "testing"

func TestHello(t *testing.T ) {
	t.Run("say Hello to people", func(t *testing.T) {
		result := Hello("Chris")
		expected := "Salve, Chris"
		
		assertCorrectMessage(t, result, expected)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		result := Hello("")
		expected := "Salve, World"

		assertCorrectMessage(t, result, expected)
	})
}

func assertCorrectMessage(t *testing.T, result, expected string) {
	t.Helper()
	if result != expected {
		t.Errorf("got %q want %q", result, expected)
	}
}
