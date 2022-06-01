package tools

import (
	"fmt"
	"time"
)

func MeasureTime(t time.Time) {
	duration := time.Since(t).Milliseconds()
	fmt.Printf("time elapsed: %vms\n", duration)
}
