package main

/*
#include <stdlib.h>

struct S {
	int a, b, c;
};

struct S *s;

void set(struct S *n) {
	s = n;
}

int getA() {
	return s->a;
}
*/
import "C"

import (
	"fmt"
	"runtime"
)

// Build and run current example.
//go:generate make run
func main() {
	set()
	gc()
	print()
}

func set() {
	cMem := C.calloc(1, C.sizeof_struct_S)
	p := (*C.struct_S)(cMem)
	p.a, p.b, p.c = 1, 2, 3
	C.set(p)
}

func print() {
	fmt.Printf("s.a = %d\n", int(C.getA()))
}

// Trigger garbage collector collects unused memory.
func gc() {
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)
	fmt.Printf("%#v\n", stats)

	runtime.GC()

	runtime.ReadMemStats(stats)
	fmt.Printf("%#v\n", stats)
}
