package main

import "C"
import "unsafe"

type Arena []unsafe.Pointer

// Alloc allocates memory on C side.
// Pointer on C memory is stored and will be free in future.
func (a *Arena) Alloc(n, size int) unsafe.Pointer {
	ptr := C.calloc(C.size_t(n), C.size_t(size))
	*a = append(*a, ptr)
	return ptr
}

// Copy copies bytes from Go to C side.
// Pointer on C memory is stored and will be free in future.
func (a *Arena) Copy(src []byte) unsafe.Pointer {
	cMem := a.Alloc(len(src), C.sizeof_char)
	return C.memmove(cMem, unsafe.Pointer(&src[0]), C.size_t(len(src))*C.sizeof_char)
}

// Free frees all C memory stored pointers.
func (a *Arena) Free() {
	for _, it := range *a {
		C.free(it)
	}
	*a = nil
}
