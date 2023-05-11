package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello world when empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello in Spannish", func(t *testing.T) {
		got := Hello("Jose", "Spanish")
		want := "Hola, Jose"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello in French", func(t *testing.T) {
		got := Hello("Jaques", "French")
		want := "Bonjour, Jaques"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello in Serbian", func(t *testing.T) {
		got := Hello("Živorad", "Serbian")
		want := "Zdravo živo, Živorad"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
