package main

import "fmt"

func squares() func() int {

	var x int

	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()

	fmt.Printf("%d\n", f())
	fmt.Printf("%d\n", f())
	fmt.Printf("%d\n", f())
	fmt.Printf("%d\n", f())
	fmt.Printf("%d\n", f())
}
