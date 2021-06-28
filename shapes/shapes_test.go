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
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{shape: Circle{Radius: 10.0}, hasArea: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		want := tt.hasArea
		if got != want {
			t.Errorf("%#v - want %g, but got %g", tt.shape, want, got)
		}
	}

}
