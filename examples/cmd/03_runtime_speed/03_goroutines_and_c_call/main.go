package main

/*
#include <unistd.h>

static inline void call() {
	sleep(1);
}
*/
import "C"

import (
	"sync"
)

// Build and run current example.
//go:generate make run-less
func main() {
	const N = 10_000
	wg := sync.WaitGroup{}

	for i := 0; i < N; i++ {
		wg.Add(1)
		go func() {
			C.call()
			wg.Done()
		}()
	}

	wg.Wait()
}
