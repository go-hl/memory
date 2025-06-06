package memory

import (
	"context"
	"fmt"
	"time"
)

// Peak store the max memory peak using your checkers.
type Peak struct {
	Alloc uint64
	Sys   uint64
}

func (p *Peak) update(alloc, sys uint64) {
	if alloc > p.Alloc {
		p.Alloc = alloc
	}

	if sys > p.Sys {
		p.Sys = sys
	}
}

// CheckSleep stay checking memory stats in a simple infinite loop.
// Works like [CheckTicker].
func (p *Peak) CheckSleep(delay time.Duration, prints ...Prints) context.CancelFunc {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				alloc, _, sys, _, _ := Stats(prints...)
				p.update(alloc, sys)
				time.Sleep(delay)
			}
		}
	}()

	return cancel
}

// CheckTicker stay checking memory stats in a for select with *time.Ticker.
// Works like [CheckSleep].
func (p *Peak) CheckTicker(delay time.Duration, prints ...Prints) context.CancelFunc {
	ctx, cancel := context.WithCancel(context.Background())
	ticker := time.NewTicker(delay)

	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				alloc, _, sys, _, _ := Stats(prints...)
				p.update(alloc, sys)
			}
		}
	}()

	return cancel
}

// Print prints memory peak stats with all other available stats.
func (p Peak) Print(prints ...Prints) {
	if len(prints) > 0 {
		prints[0].Alloc = true
		prints[0].Sys = true
	} else {
		prints = []Prints{{Alloc: true, Sys: true}}
	}
	_, total, _, gcf, gc := mem(&Mem)
	print(peak, p.Alloc, total, p.Sys, gcf, gc, prints...)
}

func (p Peak) String() string {
	return fmt.Sprintf(
		"memory peak: alloc %dMiB / sys %dMiB",
		p.Alloc/1024/1024, p.Sys/1024/1024,
	)
}
