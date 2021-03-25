package main

/*
#include <stdint.h>
*/
import "C"
import "unsafe"

//export go_reverse
func go_reverse(data *C.uint8_t, length C.size_t, output *C.uint8_t) {
	// Convert directly C memory into Go array without copy.
	goData := (*[256]byte)(unsafe.Pointer(data))

	// Convert directly C memory into Go array without copy.
	goOutput := (*[256]byte)(unsafe.Pointer(output))

	var i C.size_t
	for ; i < length; i++ {
		goOutput[length-1-i] = goData[i]
	}
}

// Build C shared library.
// go build -buildmode=c-shared
//go:generate make build-c-shared-lib
func main() {}
