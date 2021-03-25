package main

/*
#include <stdint.h> // for uint8_t type
#include <stdio.h>

typedef struct {
	uint8_t *data;
	size_t len;
} S;

void process(S *s) {
	for (int i = 0; i < s->len; ++i) {
		printf("%d\n", s->data[i]);
	}
}
*/
import "C"

// Build and run current example.
//go:generate make run
func main() {
	data := []byte{1, 2, 3}
	process(data)
}

func process(data []byte) {
	a := Arena{}
	defer a.Free()

	s := (*C.S)(a.Alloc(1, int(C.sizeof_S)))
	s.data = (*C.uint8_t)(a.Copy(data))
	s.len = C.size_t(len(data))

	C.process(s)
}
