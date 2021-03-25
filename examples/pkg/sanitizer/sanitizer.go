package sanitizer

// Minimum required GCC version 4.8

// Helpful links
// https://gcc.gnu.org/onlinedocs/gcc/Instrumentation-Options.html

// #cgo linux CFLAGS: -fsanitize=address -fsanitize=leak
// #cgo linux LDFLAGS: -fsanitize=address -fsanitize=leak
// void __lsan_do_leak_check();
import "C"

func SanitizeMemory() {
	C.__lsan_do_leak_check()
}
