package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
打印文件/输入的字符串次数
*/
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dumFile: %v\n", err)
				continue
			}

			countLines(f, counts)
			f.Close()
		}
	}

	for word, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d\n", word, n)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	var n int
	for input.Scan() {
		if n >= 10 {
			break
		}
		counts[input.Text()]++
		n++
	}
}
