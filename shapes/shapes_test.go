package main

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rectangle{10.0, 10.0}
	want := 40.0
	got := Perimeter(r)
	if got != want {
		t.Errorf("want %g, but got %g", want, got)
	}

}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {

		got := shape.Area()
		if got != want {
			t.Errorf("want %g, but got %g", want, got)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		r := Rectangle{10.0, 10.0}
		want := 100.0
		checkArea(t, r, want)

	})

	t.Run("circles", func(t *testing.T) {
		c := Circle{10.0}
		want := 314.1592653589793
		checkArea(t, c, want)
	})

}
