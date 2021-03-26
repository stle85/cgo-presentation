#include <stdint.h>
#include <libhash.h>
#include "time.h"

// Linux
// Time measured: 0.342 seconds
int main(const int argc, const char **argv) {
    const int N = 2000000;
    const int length = 16;
    uint8_t data[length];
    uint8_t output[length];

    struct timeval begin = get_time();

    for (int i = 0; i < N; ++i) {
        md5sum(data, length, output);
    }

    struct timeval end = get_time();
    measure_time(begin, end);

    return 0;
}
