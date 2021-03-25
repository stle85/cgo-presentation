package main

// Minimum required GCC version 4.8

import "C"

/*
#cgo CFLAGS: -I3rdparty/target/build/include
#cgo LDFLAGS: 3rdparty/target/build/lib/libfoo.a
#include <foo.h>
*/
import "C"
import "examples/pkg/sanitizer"

// Build and run current example.
//go:generate make run
func main() {
	defer sanitizer.SanitizeMemory()

	const N = 1_000
	C.foo(C.int(N))
}
