package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
打印输入的字符串次数大于1的字符串
*/
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	n := 0
	for input.Scan() {
		if n >= 10 {
			break
		}
		counts[input.Text()]++
		n++
		for k, v := range counts {
			println(k, v)
		}
	}

	for k, v := range counts {
		if v > 1 {
			fmt.Printf("%s\t%d\n", k, v)
		}
	}
}
