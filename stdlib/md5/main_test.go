package main

import "testing"

func BenchmarkMD5_1(b *testing.B) {
	b.N = 10000000
	for i := 0; i < b.N; i++ {
		Md51("12345678901234567890")
	}
}
func BenchmarkMD5_2(b *testing.B) {
	b.N = 10000000
	for i := 0; i < b.N; i++ {
		Md52("12345678901234567890")
	}
}
func BenchmarkMD5_3(b *testing.B) {
	b.N = 10000000
	for i := 0; i < b.N; i++ {
		Md53("12345678901234567890")
	}
}
func BenchmarkMD5_4(b *testing.B) {
	b.N = 10000000
	for i := 0; i < b.N; i++ {
		Md54("12345678901234567890")
	}
}

func BenchmarkSha1(b *testing.B) {
	b.N = 10000000
	for i := 0; i < b.N; i++ {
		Sha1("12345678901234567890")
	}
}
