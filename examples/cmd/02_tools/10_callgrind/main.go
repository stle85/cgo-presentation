package main

// You need to install valgrind on your machine.
// Command for Ubuntu: sudo apt install -y valgrind

/*
#cgo LDFLAGS: -lm
#include <math.h>

double foo() {
	double _ = 0;
	for (int j = 0; j < 2e4; ++j) {
		_ += pow(j, 7);
	}
	return _;
}
*/
import "C"

// Build and run current example under callgrind.
//go:generate make callgrind

// Show the last saved callgrind report.
//go:generate make show-callgrind-report

func main() {
	C.foo()
}
