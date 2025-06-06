package memory

import (
	"fmt"
	"runtime"

	// docs
	_ "log"
)

// Printer can change to [log.Printf] or another as desired.
var Printer func(format string, args ...any) = func(format string, args ...any) {
	fmt.Printf(format, args...)
}

// Mem is the global instance of [runtime.MemStats] that contains memory stats.
var Mem runtime.MemStats
