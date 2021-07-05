package main

import "fmt"

func double(x int) int {
	return x + x
}

func triple(x int) (result int) {
	defer func() {
		result += x
	}()

	return double(x)
}

func main() {
	fmt.Println(triple(5))
}