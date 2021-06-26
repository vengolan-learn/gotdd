package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Charan")
	want := "Hello, Charan"
	if got != want {
		t.Errorf("got %q want %q", got, want)

	} else {
		fmt.Println(Hello("Tester"))
	}
}
