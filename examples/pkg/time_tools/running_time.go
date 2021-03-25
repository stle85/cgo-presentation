package time_tools

import (
	"testing"
	"time"
)

func RunningTime() time.Time {
	return time.Now()
}

func Track(startTime time.Time, t *testing.T) {
	endTime := time.Now()
	t.Logf("=>  Took %v", endTime.Sub(startTime))
}
