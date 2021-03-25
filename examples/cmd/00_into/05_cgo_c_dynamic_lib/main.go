package main

import "C"

//export cgo_sum
func cgo_sum(a, b C.int) C.int {
	return a + b
}

//export cgo_div
func cgo_div(a, b C.int) (C.int, *C.char) {
	if b == 0 {
		return 0, C.CString("divide on zero")
	}
	return a / b, nil
}

// Build C shared library from this example.
// go build -buildmode=c-shared
//go:generate make build-c-shared-lib
func main() {}
