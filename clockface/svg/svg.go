package svg

import (
	"fmt"
	"io"
	"time"

	cf "github.com/vengolan-learn/gotdd/clockface"
)

const (
	clockCenterX = 150
	clockCenterY = 150

	clockSecondHandLength = 90
	clockMinuteHandLength = 80
	clockHourHandLength   = 50
)

//writes an SVG represenation of an analguge clock
func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	SecondHand(w, t)
	MinuteHand(w, t)
	HourHand(w, t)
	io.WriteString(w, svgEnd)
}

func makeHand(p cf.Point, length float64) cf.Point {
	p = cf.Point{X: p.X * length, Y: p.Y * length}
	p = cf.Point{X: p.X, Y: -p.Y}
	p = cf.Point{X: p.X + clockCenterX, Y: p.Y + clockCenterY}
	return p
}

func HourHand(w io.Writer, t time.Time) {
	p := cf.HoursHandPoint(t)
	p = makeHand(p, clockHourHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func MinuteHand(w io.Writer, t time.Time) {
	p := cf.MinuteHandPoint(t)
	p = makeHand(p, clockMinuteHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func SecondHand(w io.Writer, tm time.Time) {
	p := cf.SecondHandPoint(tm)
	p = makeHand(p, clockSecondHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
