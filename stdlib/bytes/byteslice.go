package main

import "bytes"

func byteSlice() {
	s := []byte{'a', 'b', 'c', 1, 2, 3, 2, 2}

	println(bytes.Contains(s, []byte{'a'}))
	println(bytes.Contains(s, []byte{1, 3}))

	println(bytes.Count(s, []byte{2, 3}))
	println(bytes.Count(s, []byte{2}))
}
