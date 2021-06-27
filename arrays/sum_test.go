package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 7, 5, 9, 4, 1, 2, 3}
		sum := Sum(numbers)
		expected := 32
		if sum != expected {
			t.Errorf("got  %d but exepcted %d given %v", sum, expected, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	numbers := [][]int{{1, 4}, {3, 6}}
	want := []int{5, 9}
	got := SumAll(numbers...)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wnat %v, but got %v", want, got)
	}
}
