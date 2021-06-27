package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	}

	t.Run("saying Hello to people", func(t *testing.T) {
		got := Hello("Charan")
		want := "Hello, Charan"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello when an empty string is passed", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

}
