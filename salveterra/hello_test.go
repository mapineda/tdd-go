package main

import "testing"

func TestHello(t *testing.T ) {
	t.Run("say Hello to people", func(t *testing.T) {
		result := Greet("Chris", "")
		expected := "Salve, Chris"
		assertCorrectMessage(t, result, expected)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		result := Greet("", "")
		expected := "Salve, World"
		assertCorrectMessage(t, result, expected)
	})

	t.Run("in Spanish", func(t *testing.T) {
		result := Greet("Elodie", "Spanish")
		expected := "Hola, Elodie"
		assertCorrectMessage(t, result, expected)
	})

	t.Run("In English", func(t *testing.T) {
		result := Greet("Chris", "English")
		expected := "Hello, Chris"
		assertCorrectMessage(t, result, expected)
	})

	t.Run("In French", func(t *testing.T) {
		result := Greet("Chris", "French")
		expected := "Bonjour, Chris"
		assertCorrectMessage(t, result, expected)
	})
}

func assertCorrectMessage(t *testing.T, result, expected string) {
	t.Helper()
	if result != expected {
		t.Errorf("got %q want %q", result, expected)
	}
}
