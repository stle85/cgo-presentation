#pragma once

#include <stdio.h>
#include <sys/time.h>
#include <unistd.h>

static inline struct timeval get_time() {
    struct timeval tm;
    gettimeofday(&tm, 0);
    return tm;
}

static inline void measure_time(struct timeval begin, struct timeval end) {
    const long seconds = end.tv_sec - begin.tv_sec;
    const long microseconds = end.tv_usec - begin.tv_usec;
    const double elapsed = seconds + microseconds*1e-6;
    printf("Time measured: %.3f seconds.\n", elapsed);
}