package main

/*
#include <stdint.h> // for uint8_t type
#include <stdlib.h> // for calloc and free funcs
#include <string.h> // for memmove func
#include <stdio.h>

typedef struct {
	uint8_t *data;
	size_t len;
} S;

void process(S *s) {
	for (int i = 0; i < s->len; ++i) {
		printf("%d\n", s->data[i]);
	}
}
*/
import "C"

import (
	"unsafe"
)

// Build and run current example.
//go:generate make run
func main() {
	data := []byte{1, 2, 3}
	process(data)
}

func process(data []byte) {
	// Alloc C structure
	s := (*C.S)(C.calloc(1, C.sizeof_S))
	defer C.free(unsafe.Pointer(s))

	// Store byte slice length
	s.len = C.size_t(len(data))

	// Alloc memory for byte from data on C side
	cMem := C.calloc(s.len, C.sizeof_char)
	defer C.free(cMem)

	// Copy memory from Go to C
	s.data = (*C.uint8_t)(C.memmove(cMem, unsafe.Pointer(&data[0]), C.sizeof_char*s.len))

	C.process(s)
}
