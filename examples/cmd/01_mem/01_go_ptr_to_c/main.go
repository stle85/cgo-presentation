package main

/*
struct S {
	int a, b, c;
};

struct S *s;

void set(struct S *n) {
	s = n;
}

int get() {
	return s->a;
}
*/
import "C"

import (
	"fmt"
)

// Build and run current example.
//go:generate make run
func main() {
	set()
	print()
}

func set() {
	p := &C.struct_S{a: 1, b: 2, c: 3}
	C.set(p)
}

func print() {
	fmt.Printf("s.a = %d\n", int(C.get()))
}
