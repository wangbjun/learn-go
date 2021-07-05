package main

import (
	"container/heap"
	"fmt"
)

func main() {
	s := students{
		{
			name:  "alan",
			score: 90},
		{
			name:  "jun",
			score: 80},
		{
			name:  "jake",
			score: 85},
	}
	//
	//fmt.Printf("%v\n", sort.Reverse(s))
	//
	//fmt.Printf("%v\n", sort.IsSorted(s))
	//
	//sort.Sort(s)
	//
	//fmt.Printf("%v\n", sort.IsSorted(s))
	//fmt.Printf("%v\n", s)
	//
	//strings := []string{"abc", "cd", "bhd", "lud"}
	//
	//sort.Strings(strings)
	//
	//fmt.Printf("%v\n", strings)
	//
	//searchStrings := sort.SearchStrings(strings, "cd")
	//
	//println(searchStrings)

	heap.Init(&s)
	heap.Push(&s, student{
		name:  "abc",
		score: 60,
	})
	heap.Push(&s, student{
		name:  "gh",
		score: 99,
	})
	//heap.Pop(&s)
	fmt.Printf("%v\n", s)
}
