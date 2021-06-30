package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++

}

func TestCountdown(t *testing.T) {

	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	Countdown(buffer, spySleeper)
	got := buffer.String()
	want := `3
2
1
Go!`
	if got != want {
		t.Errorf("want %q but got %q", want, got)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, wnat 4, but got %d", spySleeper.Calls)
	}
}
