package main

/*
//Pre define our callback to access to it in Go via 'C' pseudo-package.
void callback();

// 'static' will hide this function from export functions table.
static inline void processCallback(void (f)(void *), void *data) {
	f(data);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// Build and run current example.
//go:generate make run
func main() {
	process()
}

// Global map for storing Go callback data.
var handles = make(map[uintptr]interface{})

func process() {
	// Store data for this callback in global map.
	handles[42] = "test"

	// Call C functions with our callback function and index on our data.
	C.processCallback(
		(*[0]byte)(C.callback),
		unsafe.Pointer(uintptr(42)),
	)
}

//export callback
func callback(ptr unsafe.Pointer) {
	// Cast unsafe.Pointer to index for global map.
	data := handles[uintptr(ptr)]

	fmt.Printf("data = %#v\n", data)
}
