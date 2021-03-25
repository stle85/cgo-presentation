package main

import "C"
import (
	"fmt"

	"examples/pkg/profiler"
)

// Build and run current example.
//go:generate make run
func main() {

	// Run Go profiler to check heap allocations.
	defer profiler.MustStart(
		profiler.ProfilePath("."),
		profiler.GenerateSvg(),
		profiler.MemProfileAllocs(),
	).Stop()

	const N = 100_000

	for i := 0; i < N; i++ {
		// CString allocates memory on C side and copy string from Go to C.
		C.CString(
			fmt.Sprintf(
				"Lorem Ipsum is simply dummy text of the printing and typesetting industry. %d",
				i,
			),
		)
	}
}
