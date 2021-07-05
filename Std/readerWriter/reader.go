package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	bytes, _ := readFrom(os.Stdin, 10)
	println(string(bytes))

	from, _ := readFrom(strings.NewReader("111111111111"), 12)
	println(string(from))
}

func readFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}
