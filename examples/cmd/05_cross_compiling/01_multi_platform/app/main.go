package main

// You need to install mingw tool chain on your machine.
// Command for Ubuntu: sudo apt install -y gcc-mingw-w64-x86-64

// Helpful links
//
// List of GOOS and GOARCH
// https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63

/*
#cgo CFLAGS: -I${SRCDIR}/../lib
#cgo linux,amd64 LDFLAGS: ${SRCDIR}/../lib/libsum_glibc_x64.a
#cgo windows !amd64 LDFLAGS: ${SRCDIR}/../lib/libsum_mingw_x32.a
#include <libsum.h>
*/
import "C"
import "fmt"

// Build application.
//go:generate make build

// Build application for Windows Platform.
//go:generate make build -e TARGET=app.exe -e CC=/usr/bin/x86_64-w64-mingw32-gcc-win32 -e CGO_ENABLED=1 -e GOOS=windows

// Build and application.
//go:generate make run
func main() {
	result := C.sum(10, 20)
	fmt.Println(result)
}
