package main

// Need to install libpng on your machine
// Also your machine should have pkg-config
// About pkg-config read more on wikipedia https://en.wikipedia.org/wiki/Pkg-config

// For Ubuntu:  sudo apt install -y libpng-dev

/*
#cgo pkg-config: libpng
#include <png.h>
*/
import "C"

// build and run current example
//go:generate make run
func main() {}
