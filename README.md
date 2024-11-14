# Memory

Come capabilities to log and get some program and system memory stats.

## Usage

Example:

```go
package main

import (
	"log"
	"memory/memory"
	"runtime"
	"time"
)

func main() {
	var (
		mp memory.Peak
		// cancel context.CancelFunc
	)
	memory.Printer = log.Printf

	// go func() {
	// 	cancel = mp.CheckWithSleep(time.Second)
	// }()
	go mp.CheckWithSleep(time.Second, true)

	var allocer [][]any
	for index := 0; index < 10; index++ {
		allocer = append(allocer, make([]any, 999999))
		if index == 7 {
			allocer = nil
			runtime.GC()
			// cancel()
		}
		time.Sleep(time.Millisecond * 500)
	}

	runtime.GC()
	memory.CheckStats()
	log.Println(mp)
}
```
