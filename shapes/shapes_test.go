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

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10.0}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		want := tt.want
		if got != want {
			t.Errorf("want %g, but got %g", want, got)
		}
	}

}
