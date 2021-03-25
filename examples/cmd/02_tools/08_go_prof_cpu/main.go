package main

/*
#cgo LDFLAGS: -lm
#include <math.h>

double foo() {
	double _ = 0;
	for (int j = 0; j < 2e8; ++j) {
		_ += pow(j, 7);
	}
	return _;
}
*/
import "C"
import "examples/pkg/profiler"

// Build and run current example.
//go:generate make run
func main() {

	// Run Go profiler to check application CPU calls.
	defer profiler.MustStart(
		profiler.ProfilePath("."),
		profiler.GenerateSvg(),
		profiler.CPUProfile(),
	).Stop()

	C.foo()
}
