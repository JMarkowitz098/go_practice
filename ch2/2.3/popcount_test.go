package main


import (
	"testing"

	"./popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(num)
	}
}

func BenchmarkWithLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.WithLoop(num)
	}
}

