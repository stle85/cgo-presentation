package main

// Need to install libpng on your machine.

// For Ubuntu:  sudo apt install -y libpng-dev

/*
#cgo CFLAGS: -DPNG_DEBUG=1
#cgo amd64 386 CFLAGS: -DX86=1
#cgo LDFLAGS: -lpng
#include <png.h>
*/
import "C"

// Build and run current example.
//go:generate make run
func main() {}
