package main

/*
#cgo linux LDFLAGS: -ldl
#include <dlfcn.h>
#include <limits.h>
#include <stdlib.h>
#include <stdint.h>

static uintptr_t openLib(const char* path, char** err) {
	void* h = dlopen(path, RTLD_LAZY);
	if (h == NULL) {
		*err = (char*)dlerror();
	}
	return (uintptr_t)h;
}

static int closeLib(uintptr_t h) {
    if (h == 0) {
       return 0;
    }
	return dlclose((void*)h);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type LibHandler C.uintptr_t

func OpenLib(filePath string) (LibHandler, error) {
	cPath := make([]byte, C.PATH_MAX+1)
	cRelName := make([]byte, len(filePath)+1)
	copy(cRelName, filePath)
	if C.realpath(
		(*C.char)(unsafe.Pointer(&cRelName[0])),
		(*C.char)(unsafe.Pointer(&cPath[0]))) == nil {
		return 0, fmt.Errorf("realpath failed")
	}

	var cErr *C.char
	h := C.openLib((*C.char)(unsafe.Pointer(&cPath[0])), &cErr)
	if h == 0 {
		return 0, fmt.Errorf("OpenLib('%s'): %s", filePath, C.GoString(cErr))
	}
	return LibHandler(h), nil
}

func CloseLib(h LibHandler) error {
	code := C.closeLib(C.uintptr_t(h))
	if code != 0 {
		return fmt.Errorf("CloseLib(): cannot close lib. code: %d", code)
	}
	return nil
}
