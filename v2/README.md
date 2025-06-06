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

func init() {
	memory.Printer = log.Printf // this is optional
}

func main() {
	var mp memory.Peak
	prints := memory.Prints{All: true}

	// can call (many times) before the parts you desire check
	// or in begin of the program
	cancel := mp.CheckSleep(time.Second, prints)

	process()

	// also can call (many times) after the parts you desire finish the check
	// or in ending of the program
	// cancel()
	// log.Println(mp)

	executions()

	cancel()        // stop collect and print stats
	log.Println(mp) // print the collect memory peak

	runtime.GC()
	memory.Stats(prints) // single print stats
}

func process() {
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
