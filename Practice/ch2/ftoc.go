package main

import "fmt"

const name = `
	what is your name? I don't know '
`

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   // "212°F = 100°C"

	fmt.Println("-----------------------------------------------------")

	charTest()
	hello()
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func charTest() {
	s := "Hello World!"
	// c := s[len(s)]
	fmt.Println(len(s))
	fmt.Printf("%c,%c\n", s[0], s[7])
	// fmt.Println(c)

	chinese := "我爱你"

	fmt.Printf("%s\n", chinese[0:3]) //我

	fmt.Println(chinese[0:3] + "是谁？") //我是谁？

	fmt.Print(name)
}

func hello() {
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}
