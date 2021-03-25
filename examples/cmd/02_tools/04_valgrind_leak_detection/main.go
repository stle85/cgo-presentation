package main

// You need to install valgrind on your machine.
// Command for Ubuntu: sudo apt install -y valgrind

// Helpful links
//
// https://github.com/golang/go/issues/27610
// It doesn't seem like we should change how the Go library behaves only to make valgrind happy.
// We need some way to tell valgrind that this particular read of uninitialized memory is OK.
// Is there any way to do that?
//
// https://www.valgrind.org/docs/manual/faq.html#faq.deflost
// * "definitely lost" means your program is leaking memory -- fix those leaks!
// * "indirectly lost" means your program is leaking memory in a pointer-based structure.
//   (E.g. if the root node of a binary tree is "definitely lost", all the children will be "indirectly lost".)
//   If you fix the "definitely lost" leaks, the "indirectly lost" leaks should go away.
// * "possibly lost" means your program is leaking memory, unless you're doing unusual things with pointers that could
//    cause them to point into the middle of an allocated block; see the user manual for some possible causes.
//    Use --show-possibly-lost=no if you don't want to see these reports.
// * "still reachable" means your program is probably ok -- it didn't free some memory it could have.
//    This is quite common and often reasonable. Don't use --show-reachable=yes if you don't want to see these reports.
// * "suppressed" means that a leak error has been suppressed.
//    There are some suppressions in the default suppression files. You can ignore suppressed errors.

import "C"

// #include <stdlib.h>
import "C"

// Build and run current example under valgrind.
//go:generate make valgrind
func main() {
	const N = 1_000

	for i := 0; i < N; i++ {
		// Allocate memory on C side
		C.calloc(1, 100)
	}
}
