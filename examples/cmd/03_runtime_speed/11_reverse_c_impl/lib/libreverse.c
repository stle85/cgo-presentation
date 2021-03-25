#include "libreverse.h"
#include <openssl/md5.h>

void reverse(const uint8_t *data, const int length, uint8_t *output) {
    for (size_t i = 0; i < length; ++i) {
        output[length - 1 - i] = data[i];
    }
}
