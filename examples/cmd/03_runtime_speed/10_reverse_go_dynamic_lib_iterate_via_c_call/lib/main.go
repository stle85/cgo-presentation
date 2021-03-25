package main

/*
#include <stdint.h>

static inline void reverse(uint8_t *data, size_t length, uint8_t *output) {
    for (size_t i = 0; i < length; ++i) {
        output[length - 1 - i] = data[i];
    }
}
*/
import "C"

//export go_reverse
func go_reverse(data *C.uint8_t, length C.size_t, output *C.uint8_t) {
	// Copy input data to output via C.
	C.reverse(data, length, output)
}

// Build C shared library.
// go build -buildmode=c-shared
//go:generate make build-c-shared-lib
func main() {}
