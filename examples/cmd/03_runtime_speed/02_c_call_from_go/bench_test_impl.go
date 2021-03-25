package main

/*
#cgo LDFLAGS: -lm
#include <math.h>

static inline double call1() {
	return pow(3, 4);
}

static inline double call2(int n) {
	double _;
	for (int i = 0; i < n; ++i) {
		_ += call1();
	}
	return	_;
}
*/
import "C"
import (
	"testing"

	"examples/pkg/time_tools"
)

const N = 1_000_000

// platform: linux
// cpu: AMD Athlon 200GE with Radeon Vega Graphics
// 83.081479ms

// Run test.
//go:generate make go-test -e TEST_NAME=TestLoopInGo
func testLoopInGo(t *testing.T) {
	defer time_tools.Track(time_tools.RunningTime(), t)

	for i := 0; i < N; i++ {
		C.call1()
	}
}

// platform: linux
// cpu: AMD Athlon 200GE with Radeon Vega Graphics
// 973.698µs

// Run test.
//go:generate make go-test -e TEST_NAME=TestLoopInC
func testLoopInC(t *testing.T) {
	defer time_tools.Track(time_tools.RunningTime(), t)

	res := C.call2(C.int(N))
	t.Logf("result = %v", res)
}
