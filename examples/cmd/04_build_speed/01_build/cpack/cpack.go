package cpack

/*
#cgo LDFLAGS: -lcrypto
#include <openssl/md5.h>

void process(int a, int b, int c, int d) {
    MD5_CTX m;
	char output[16];

    MD5_Init(&m);
    MD5_Update(&m, &a, sizeof(int));
    MD5_Update(&m, &b, sizeof(int));
    MD5_Update(&m, &c, sizeof(int));
    MD5_Update(&m, &d, sizeof(int));
    MD5_Final(output, &m);
}
*/
import "C"

type Cint C.int

func Process(a, b, c, d Cint) {
	C.process(C.int(a), C.int(b), C.int(c), C.int(d))
}
