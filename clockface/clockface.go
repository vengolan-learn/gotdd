package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

func hoursInRadians(t time.Time) float64 {

	return (minutesInRadians(t) / 12) + (math.Pi/6)*float64(t.Hour()%12)
}

func HoursHandPoint(t time.Time) Point {
	rad := hoursInRadians(t)
	return angleToPoint(rad)
}

func minutesInRadians(tm time.Time) float64 {
	return (secondsInRadians(tm) / 60) +
		(math.Pi/30)*float64(tm.Minute())
}

func MinuteHandPoint(tm time.Time) Point {
	rad := minutesInRadians(tm)
	return angleToPoint(rad)
}

func secondsInRadians(tm time.Time) float64 {
	return (math.Pi / 30) * float64(tm.Second())
}

func SecondHandPoint(tm time.Time) Point {
	rad := secondsInRadians(tm)
	return angleToPoint(rad)
}
