package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Lilly", "English")
		want := "Hello Lilly"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello world when an empty string is supposed", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hola to people", func(t *testing.T) {
		got := Hello("Lilly", "Spanish")
		want := "Hola Lilly"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying Bonjour to people", func(t *testing.T) {
		got := Hello("Lilly", "French")
		want := "Bonjour Lilly"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
