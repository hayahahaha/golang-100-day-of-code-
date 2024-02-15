package main

import "testing"

func TestHelo(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("haya", "")
		want := "Hello, haya"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello to people in spanish", func(t *testing.T) {
		got := Hello("haya", "spanish")
		want := "Hola, haya"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello world in spanish", func(t *testing.T) {
		got := Hello("", "spanish")
		want := "Hola, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q ", got, want)
	}

}
