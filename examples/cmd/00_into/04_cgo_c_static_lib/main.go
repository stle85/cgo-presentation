package main

import "C"

//export cgo_sum
func cgo_sum(a, b C.int) (C.int, *C.char) {
	return a + b, nil
}

// Build C static library from this example.
// go build -buildmode=c-archive
//go:generate make build-c-static-lib
func main() {}
