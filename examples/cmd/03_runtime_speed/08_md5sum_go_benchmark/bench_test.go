package main

import (
	"crypto/md5"
	"testing"

	"examples/pkg/time_tools"
)

func hash(t *testing.T) [16]byte {
	const N = 2_000_000
	data := make([]byte, 16)
	output := [16]byte{}

	defer time_tools.Track(time_tools.RunningTime(), t)
	for i := 0; i < N; i++ {
		output = md5.Sum(data)
	}
	return output
}

// platform: linux
// cpu: AMD Athlon 200GE with Radeon Vega Graphics
// 337.851982ms

// Run test.
func TestMd5Sum(t *testing.T) {
	hash(t)
}
