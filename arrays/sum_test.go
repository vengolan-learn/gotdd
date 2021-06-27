package main

import "testing"

func Test(t *testing.T) {
	numbers := [5]int{1, 7, 5, 9, 4}
	sum := Sum(numbers)
	expected := 26
	if sum != expected {
		t.Errorf("got  %d but exepcted %d given %v", sum, expected, numbers)
	}
}
