package main

import "testing"

func BenchmarkWithoutPtr(b *testing.B) {
	benchmarkWithoutPtr(b)
}

func BenchmarkOnePtr(b *testing.B) {
	benchmarkOnePtr(b)
}

func BenchmarkTwoPtr(b *testing.B) {
	benchmarkTwoPtr(b)
}
