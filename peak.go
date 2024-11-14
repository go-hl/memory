package memory

import (
	"context"
	"fmt"
	"time"
)

// Peak store max memory peak with your checkers.
type Peak struct {
	Alloc uint64
	Sys   uint64
}

func (p Peak) String() string {
	return fmt.Sprintf(
		"memory peak: alloc %dMB - sys %dMB",
		p.Alloc/1024/1024, p.Sys/1024/1024,
	)
}

// CheckWithTicker stay checking memory stats in a for select with *time.Ticker.
func (p *Peak) CheckWithTicker(delay time.Duration, print ...bool) (cancel context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())

	ticker := time.NewTicker(delay)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				alloc, sys := CheckStats(print...)
				if alloc > p.Alloc {
					p.Alloc = alloc
				}
				if sys > p.Sys {
					p.Sys = sys
				}
			}
		}
	}()

	return
}

// CheckWithSleep stay checking memory stats in a simple infinite loop.
func (p *Peak) CheckWithSleep(delay time.Duration, print ...bool) (cancel context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				alloc, sys := CheckStats(print...)
				if alloc > p.Alloc {
					p.Alloc = alloc
				}
				if sys > p.Sys {
					p.Sys = sys
				}
				time.Sleep(delay)
			}
		}
	}()

	return
}
