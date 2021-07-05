package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)

	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	t := x / 64
	k := uint(x % 64)
	word, bit := t, k
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] = 1<<bit | s.words[word]
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)
	x.Add(7)
	x.Add(3)
	x.Add(91)
	x.Add(21)
	x.Add(33)
	x.Add(65)
	x.Add(213)
	x.Add(452)
	x.Add(252)
	x.Add(409)
	x.Add(14)
	x.Add(4)
	x.Add(9)
	x.Add(5)
	x.Add(63)
	x.Add(64)
	x.Add(65)
	x.Add(128)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String())           // "{1 9 42 144}"
	fmt.Println(x.Has(9), x.Has(123)) // "true false"
}
