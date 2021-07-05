package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

type Value interface {
	String() string
	Set(string) error
}

func main() {
	flag.Parse()

	fmt.Printf("Sleeping for %v...\n", *period)

	time.Sleep(*period)

	println()
}
