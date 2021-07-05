package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		_, r := div(a, b)
		return r, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func div(a, b int) (q, r int) {
	return a % b, a / b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args"+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func plus(a, b int) int {
	return a + b
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error")
	} else {
		fmt.Println(result)
	}
	a, b := div(13, 3)
	fmt.Println(a, b)

	result := apply(plus, 3, 4)

	fmt.Println(result)

	x, y := 3, 4
	swap(&x, &y)
	fmt.Println(x, y)
}
