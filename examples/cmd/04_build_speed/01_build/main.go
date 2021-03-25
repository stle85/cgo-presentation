package main

import "C"
import "examples/cmd/04_build_speed/01_build/cpack"

// Build time: 0:00.82

// Build application and show build time.
//go:generate make build-time
func main() {
	var a, b, c, d cpack.Cint
	cpack.Process(a, b, c, d)
}
