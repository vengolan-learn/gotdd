package main

import "math"

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Height + rect.Width)
}

func main() {

}
