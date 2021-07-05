package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Path []Point

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		fmt.Printf("i = %v\n", i)
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}

	return sum
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point

	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i := range path {
		path[i] = op(path[i], offset)
	}
}

func main() {
	p := Point{1, 2}
	q := Point{5, 6}

	fmt.Printf("%f\n", Distance(p, q))
	fmt.Printf("%f\n", p.Distance(q))

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())

	distanceFromP := p.Distance

	fmt.Printf("%f\n", distanceFromP(q))

	path := Path{
		Point{1, 2},
		Point{5, 6},
	}

	path.TranslateBy(p, true)

	fmt.Printf("%v\n", path)
}
