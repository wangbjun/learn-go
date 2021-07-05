package main

import (
	"bufio"
	"os"
	"strconv"
)

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fib2(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	var arr [20]int
	arr[0] = 0
	arr[1] = 1
	for i := 2; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}

func fibGen() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < 20; i++ {
		writer.WriteString(strconv.Itoa(fib2(i)) + "\n")
	}
	f := fibGen()
	for i := 0; i < 20; i++ {
		writer.WriteString(strconv.Itoa(f()) + "\n")
	}
}

func main() {
	writeFile("fib.txt")
}
