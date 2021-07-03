package main

import "testing"

func TestWalk(t *testing.T) {

	expected := "Rama"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if got[0] != expected {
		t.Errorf("want %v, but got %v", expected, got[0])
	}

}
