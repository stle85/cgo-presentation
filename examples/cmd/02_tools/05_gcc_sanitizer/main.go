package main

// Minimum required GCC version 4.8

import "C"

// #include <stdlib.h>
import "C"
import "examples/pkg/sanitizer"

// Build and run current example.
//go:generate make run
func main() {
	// Run GCC sanitizer.
	defer sanitizer.SanitizeMemory()

	const N = 1_000

	for i := 0; i < N; i++ {
		// Build and run current example under valgrind.
		C.calloc(1, 100)
	}
}
