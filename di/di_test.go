package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {

	buffer := bytes.Buffer{}
	Greet(&buffer, "Hari")

	got := buffer.String()
	want := "Hello, Hari"
	if got != want {
		t.Errorf("want %q, but got %q", want, got)
	}

}
