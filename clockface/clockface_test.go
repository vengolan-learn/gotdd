package clockface

import (
	"math"
	"testing"
	"time"
)

func simpleTime(hour int, minute int, second int) time.Time {
	return time.Date(312, time.October, 35, hour, minute, second, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(9, 0, 0), (math.Pi / 6) * float64(9)},
		{simpleTime(0, 1, 30), (math.Pi / (6 * 60 * 60)) * float64(90)},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := hoursInRadians(c.time)
			if !roughlyEqualFloats(got, c.angle) {
				t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
			}
		})
	}
}
func TestHourHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(6, 0, 0), Point{0, -1}},
		{simpleTime(21, 0, 0), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := HoursHandPoint(c.time)
			if !roughlyEqualPoints(got, c.point) {
				t.Fatalf("Wanted %v Point, but got %v", c.point, got)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 30, 0), Point{0, -1}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := MinuteHandPoint(c.time)
			if !roughlyEqualPoints(got, c.point) {
				t.Fatalf("Wanted %v point, but got %v", c.point, got)
			}
		})
	}

}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := minutesInRadians(c.time)
			if !roughlyEqualFloats(got, c.angle) {
				t.Fatalf("wanted %v radians, but got %v", c.angle, got)
			}
		})
	}
}

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), (math.Pi / 30) * float64(30)},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 30) * float64(45)},
		{simpleTime(0, 0, 7), (math.Pi / 30) * float64(7.0)},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondsInRadians(c.time)
			if got != c.angle {
				t.Fatalf("want %v radians, but got %v", c.angle, got)
			}

		})

	}
}

func TestSecondHandVector(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
		{simpleTime(0, 0, 0), Point{0, 1}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := SecondHandPoint(c.time)
			if !roughlyEqualPoints(got, c.point) {
				t.Fatalf("Wanted %v Point, but got %v", c.point, got)
			}
		})
	}
}

func roughlyEqualFloats(a, b float64) bool {
	const equalityTresh = 1e-6
	return math.Abs(a-b) < equalityTresh
}

func roughlyEqualPoints(a, b Point) bool {
	return roughlyEqualFloats(a.X, b.X) &&
		roughlyEqualFloats(a.Y, b.Y)
}
