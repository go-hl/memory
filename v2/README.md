# Memory

Come capabilities to log and get some program and system memory stats.

## Install

Run:
```
go get -u "github.com/go-hl/memory/v2"
```

## Usage

Example:
```go
package main

import (
	"log"
	"math"
	"math/rand"
	"runtime"
	"strconv"
	"time"

	"github.com/go-hl/memory/v2"
)

func usingPeak() {
	var mp memory.Peak
	prints := memory.Prints{All: true}
	cancel := mp.CheckSleep(time.Second, prints) // start check memory peak

	processes()
	executions()

	cancel() // stop check memory peak
	runtime.GC()
	memory.Stats(prints) // single print memory stats
	log.Println(mp)      // print the collected memory peak
}

func usingMonitor() {
	var mm memory.Monitor
	prints := memory.Prints{All: true}

	mm.Register("processes") // begin monitor memory for this one
	processes()
	mm.Shut("processes") // end monitor memory for this one

	mm.Register("executions")
	executions()
	mm.Shut("executions")

	runtime.GC()
	memory.Stats(prints) // single print memory stats
	mm.Log()             // print the monitored memory
}

func init() {
	memory.Printer = log.Printf // this is optional
}

func main() {
	usingPeak()
	runtime.GC()
	usingMonitor()
}

func processes() {
	log.Println("processing anything")

	alloc := make(map[string]int)

	for index := range 10 {
		for range 100000 {
			value := rand.Int()
			key := strconv.Itoa(value)
			alloc[key] = value
		}

		if index == 7 {
			alloc = map[string]int{}
			runtime.GC()
		}

		time.Sleep(time.Second)
	}
}

func executions() {
	log.Println("another executions")

	var alloc1 []uint64
	var alloc2 []uint64

	for range 5 {
		for range 1000000 {
			alloc1 = append(alloc1, math.MaxUint64)
			alloc2 = append(alloc2, math.MaxUint64)
		}
		time.Sleep(time.Second)
	}
}
```
