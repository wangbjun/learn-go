package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}

	return g
}

func convertToBin(n int) {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}

	fmt.Println(result)
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	const filename = "abc.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%s\n", content)
	}

	if content, err := ioutil.ReadFile(filename); err != nil {
		panic(err)
	} else {
		fmt.Printf("%s\n", content)
	}

	fmt.Println(grade(70))
	fmt.Println(grade(85))
	fmt.Println(grade(95))

	convertToBin(5)
	convertToBin(13)
	convertToBin(50)

	printFile(filename)
}
