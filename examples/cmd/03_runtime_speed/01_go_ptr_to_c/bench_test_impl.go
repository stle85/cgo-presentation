package main

/*
typedef struct {
	int a, b;
} S1;

static inline void call1(S1 *s) {}


typedef struct {
	int *a, b;
} S2;

static inline void call2(S2 *s) {}


typedef struct {
	int *a, *b;
} S3;

static inline void call3(S3 *s) {}
*/
import "C"
import "testing"

// platform: linux
// cpu: AMD Athlon 200GE with Radeon Vega Graphics
// BenchmarkWithoutPtr   68.41 ns/op

// Run benchmark test.
//go:generate make go-benchmark -e TEST_NAME=BenchmarkWithoutPtr
func benchmarkWithoutPtr(b *testing.B) {
	s := &C.S1{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C.call1(s)
	}
}

// platform: linux
// cpu: AMD Athlon 200GE with Radeon Vega Graphics
// BenchmarkOnePtr   99.59 ns/op

// Run benchmark test.
//go:generate make go-benchmark -e TEST_NAME=BenchmarkOnePtr
func benchmarkOnePtr(b *testing.B) {
	s := &C.S2{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C.call2(s)
	}
}

// platform: linux
// cpu: AMD Athlon 200GE with Radeon Vega Graphics
// BenchmarkTwoPtr   102.5 ns/op

// Run benchmark test.
//go:generate make go-benchmark -e TEST_NAME=BenchmarkTwoPtr
//go:generate make go-benchmark -e TEST_NAME=BenchmarkTwoPtr -e GODEBUG=cgocheck=0
func benchmarkTwoPtr(b *testing.B) {
	s := &C.S3{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C.call3(s)
	}
}
