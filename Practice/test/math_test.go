package main

import (
	"testing"
)

func TestDiv(t *testing.T) {
	i, err := Div(4, 2)
	if err != nil {
		t.Fail()
	}
	if i != 2 {
		t.Fail()
	}

	i, err = Div(4, 0)
	if err == nil {
		t.Fail()
	}
}

func TestDiv2(t *testing.T) {
	var tests = []struct {
		a        int
		b        int
		expected int
		err      error
	}{
		{4, 2, 2, nil},
		{4, 1, 4, nil},
		{5, 2, 2, nil},
		{4, 0, 0, DivisionByZeroError},
	}

	for _, v := range tests {
		i, err := Div(v.a, v.b)
		if i != v.expected || err != v.err {
			t.Errorf("input %d, %d, expected %d, got %d", v.a, v.b, v.expected, i)
		}
	}
}
