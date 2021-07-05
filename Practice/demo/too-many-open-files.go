package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var i = 0
	for {
		open, err := os.Open("/home/jwang/Downloads/app.txt")
		if err != nil {
			panic(fmt.Sprintf("err: %s, open %d", err, i))
		}
		open.Write([]byte("abc"))
		i++
		time.Sleep(time.Millisecond * 1000)
	}
}
