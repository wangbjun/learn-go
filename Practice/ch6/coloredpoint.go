package main

import (
	"fmt"
	"image/color"
)

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func main() {
	var p ColoredPoint
	p.X = 1
	p.Y = 3
	red := color.RGBA{255, 0, 0, 255}

	p.Color = red

	var q = ColoredPoint{Point{1, 5}, red}

	fmt.Printf("%f\n", p.X)
	fmt.Printf("%d\n", q.Color.A)
}
