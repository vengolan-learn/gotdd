package main

import (
	"os"
	"time"

	svg "github.com/vengolan-learn/gotdd/clockface/svg"
)

func main() {
	t := time.Now()
	svg.SVGWriter(os.Stdout, t)
}
