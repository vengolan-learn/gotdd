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

	if len(got) != 1 {
		t.Errorf("wrong number of ufnction calls, want %d, but got %d", 1, len(got))
	}

}
