#include <stdlib.h>

void foo(int n) {
    void *p;
    for (int i = 0; i < n; ++i) {
        p = calloc(n, 1);
    }
}
