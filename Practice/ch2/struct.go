package main

func main() {

	type Point struct {
		X, Y int
	}

	// type Circle struct {
	// 	Center Point
	// 	Radius int
	// }
	//
	// type Wheel struct {
	// 	Circle Circle
	// 	Spokes int
	// }
	//
	// var w Wheel
	// w.Circle.Center.X = 1
	// w.Circle.Center.Y = 2
	//
	// println(w.Circle.Center.X)
	// println(w.Circle.Center.Y)

	type Circle struct {
		Point
		Radius int
	}

	type Wheel struct {
		Circle
		Spokes int
	}

	var w Wheel
	w.X = 1
	w.Y = 2
	println(w.X)
	println(w.Y)
}
