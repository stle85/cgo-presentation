package main

/*
#include <stdio.h>

typedef struct {
	int a, b, c;
} S;

typedef struct {
	S *x, y;
} C;

void process(C *c) {
	printf("c.x.a = %d\n", c->x->a);
	printf("c.x.b = %d\n", c->x->b);
	printf("c.x.c = %d\n", c->x->c);
}
*/
import "C"

// Build and run current example with different GODEBUG=cgo-check=<N> env variables.
//go:generate make run
//go:generate make run -e GODEBUG=cgocheck=0
//go:generate make run -e GODEBUG=cgocheck=2
func main() {
	s := &C.S{a: 1, b: 2, c: 3}
	c := &C.C{x: s}
	C.process(c)
}
