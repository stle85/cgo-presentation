package main

import "examples/cmd/04_build_speed/02_improved_build/cpack"

// Build time: 0:00.75

// Build application and show build time.
//go:generate make build-time
func main() {
	var a, b, c, d int
	cpack.Process(a, b, c, d)
}
