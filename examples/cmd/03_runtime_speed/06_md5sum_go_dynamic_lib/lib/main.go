package main

/*
#include <stdint.h>
*/
import "C"
import (
	"crypto/md5"
	"unsafe"
)

//export go_md5sum
func go_md5sum(data *C.uint8_t, length C.size_t, output *C.uint8_t) {
	const outputLength = 16
	goLength := int(length)

	// Convert directly C memory into Go slice without copy.
	goData := (*[1 << 28]byte)(unsafe.Pointer(data))[:goLength:goLength]

	// Convert directly C memory into Go array without copy.
	goOutput := (*[outputLength]byte)(unsafe.Pointer(output))

	// Calculate md5
	*goOutput = md5.Sum(goData)
}

// Build C shared library.
// go build -buildmode=c-shared
//go:generate make build-c-shared-lib
func main() {}
