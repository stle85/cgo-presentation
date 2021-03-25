#include "libhash.h"
#include <openssl/md5.h>

void md5sum(const uint8_t *data, const int length, uint8_t *output) {
    MD5_CTX c;

    MD5_Init(&c);
    MD5_Update(&c, data, length);
    MD5_Final(output, &c);
}
