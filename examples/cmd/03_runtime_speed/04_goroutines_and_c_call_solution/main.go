package main

/*
#include <unistd.h>

static inline void call() {
	sleep(1);
}
*/
import "C"

import (
	"context"
	"runtime"
	"sync"

	"golang.org/x/sync/semaphore"
)

// Build and run current example.
//go:generate make run-less
func main() {
	const N = 10_000
	wg := sync.WaitGroup{}

	ctx := context.TODO()
	sem := semaphore.NewWeighted(int64(runtime.NumCPU() + 1))

	for i := 0; i < N; i++ {
		wg.Add(1)
		go func() {
			_ = sem.Acquire(ctx, 1)
			defer sem.Release(1)

			C.call()
			wg.Done()
		}()
	}

	wg.Wait()
}
