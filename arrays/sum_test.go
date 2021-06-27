package main

import "testing"

func Test(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 7, 5, 9, 4, 1, 2, 3}
		sum := Sum(numbers)
		expected := 32
		if sum != expected {
			t.Errorf("got  %d but exepcted %d given %v", sum, expected, numbers)
		}
	})

}
