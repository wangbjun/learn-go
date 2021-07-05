package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "2jda13131as"

	index := strings.Index(str, "13")

	fmt.Printf("%d\n", index)
}
