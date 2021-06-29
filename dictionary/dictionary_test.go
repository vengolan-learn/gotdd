package main

import (
	"testing"
)

func TestDelete(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dict := Dictionary{word: definition}

	dict.Delete(word)
	_, err := dict.Search(word)
	assertError(t, err, ErrNotFound)

}

func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dict := Dictionary{word: definition}
		newDefinition := "new definition"
		dict.Update(word, newDefinition)
		assertDefinition(t, dict, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dict := Dictionary{}

		err := dict.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExists)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "test"
		definition := "this is just a test"

		dict.Add(word, definition)
		assertDefinition(t, dict, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dict := Dictionary{word: definition}
		err := dict.Add(word, "new test")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, word, definition)
	})

}

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known workd", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, want, got)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		assertError(t, err, ErrNotFound)
		if err == nil {
			t.Fatal("fexpected to get an error")
		}

	})
}

func assertDefinition(t testing.TB, dict Dictionary, word, definition string) {
	t.Helper()
	got, err := dict.Search("test")
	if err != nil {
		t.Fatal("shoud find added word:", err)
	}
	assertStrings(t, got, definition)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("want %#v, but got %#v", want, got)

	}

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("want %q, but got %q", want, got)
	}

}
