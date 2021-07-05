package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var packageValue = 35

func variableZeroValue() {
	var a int
	var s string

	fmt.Printf("%d, %q\n", a, s)
}

func variableInitValue() {
	var a, b = 3, 4
	var s = "abc"
	fmt.Println(a, b, s)
}

func variableShorter() {
	a, b, c := 3, 4, "abc"
	fmt.Println(a, b, c, packageValue)
}

func triangle() {
	var a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))

	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

func constValue() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))

	fmt.Println(filename, c)
}

func enumValue() {
	const (
		cpp = iota
		java
		golang
		php
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, java, golang, php)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello World!")
	variableZeroValue()
	variableInitValue()
	variableShorter()
	euler()
	triangle()
	constValue()
	enumValue()
}
