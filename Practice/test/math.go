package main

import "errors"

var DivisionByZeroError = errors.New("division by zero")

func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, DivisionByZeroError
	}
	return a / b, nil
}
