package main

import "testing"

func TestHello(t *testing.T) {

	verifyCorrectMessage := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result '%s', expected '%s", result, expected)
		}
	}

	t.Run("says Hello to people", func(t *testing.T) {
		result := Hello("Chris", "english")
		expected := "Hello, Chris"

		verifyCorrectMessage(t, result, expected)
	})

	t.Run("says 'Hello, World' when an empty string is passed", func(t *testing.T) {
		result := Hello("", "")
		expected := "Hello, World"

		verifyCorrectMessage(t, result, expected)
	})

	t.Run("in spanish", func(t *testing.T) {
		result := Hello("Elodie", "spanish")
		expected := "Hola, Elodie"

		verifyCorrectMessage(t, result, expected)
	})
}
