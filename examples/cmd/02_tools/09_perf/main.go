package main

// You need to install perf on your machine.
// And allow it investigate your running application.

// For Ubuntu:
// sudo apt install -y linux-tools-5.8.0-45-generic   # for my current kernel version
// sudo sh -c 'echo -1 >  /proc/sys/kernel/perf_event_paranoid'   # for running application without root rights.

/*
#cgo LDFLAGS: -lm
#include <math.h>

double foo() {
	double _ = 0;
	for (int j = 0; j < 2e8; ++j) {
		_ += pow(j, 7);
	}
	return _;
}
*/
import "C"

// Build and run current example under perf.
//go:generate make perf

// Show the last saved perf report.
//go:generate make show-perf-report

func main() {
	C.foo()
}
