package main

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"
	"testing"
)

func BenchmarkRand1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Intn(1000)
	}
}

func BenchmarkRand2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		crand.Int(crand.Reader, big.NewInt(1000))
	}
}
