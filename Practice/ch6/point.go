package main

import "fmt"

// type Point struct {
// 	X, Y float64
// }

func (p *Point) ScaleBy(factor float64) {
	fmt.Printf("%v\n", p)

	newP := Point{5, 6}

	p = &newP // 改变指针的值并不会影响其原来的值

	p.X = 10

	p.X *= factor // 改变指针元素的值
	p.Y *= factor
}

func main() {
	r := Point{1, 2}

	r.ScaleBy(2)

	fmt.Printf("%f", r)
}
