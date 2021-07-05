package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
)

func main() {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", base64.StdEncoding.EncodeToString(b))

	n, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", n)
}
