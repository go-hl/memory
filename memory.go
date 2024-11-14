package memory

import (
	"fmt"
	"runtime"
)

// Printer can change to fmt.Printf() or another as desired.
var Printer func(format string, args ...any) = func(format string, args ...any) {
	fmt.Printf(format, args...)
}

// CheckStats return atual alloc and sys memory and print this stats.
func CheckStats(print ...bool) (alloc, sys uint64) {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	if len(print) > 0 && print[0] {
		Printer(
			"alloc: %dMB - total: %dMB / sys: %dMB\n",
			stats.Alloc/1024/1024,
			stats.TotalAlloc/1024/1024,
			stats.Sys/1024/1024,
		)
	}
	return stats.Alloc, stats.Sys
}
