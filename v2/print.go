package memory

import (
	"math"
	"strings"
)

type allocs struct {
	alloc *uint64
	total *uint64
}

type gcs struct {
	gcf *uint32
	gc  *uint32
}

type params struct {
	allocs *allocs
	sys    *uint64
	gcs    *gcs
}

// Prints sets if print or not the stats.
type Prints struct {
	// All is all available stats.
	All bool

	// Allocs is an alias to [Prints.Alloc] and [Prints.Total] stats.
	Allocs bool
	// Memory is an alias to [Prints.Allocs] and [Prints.Sys] stats.
	Memory bool
	// GCs is an alias to [Prints.GCF] and [Prints.GC] stats.
	GCs bool

	// Alloc is the atual memory allocated in Heap ([runtime.MemStats.Alloc])
	Alloc bool
	// Total is the total memory allocated in Heap throughout the execution ([runtime.MemStats.TotalAlloc])
	Total bool
	// Sys is the reserved memory in OS for the programa ([runtime.MemStats.Sys])
	Sys bool
	// GCF is the forced execution cycles of GC ([runtime.MemStats.NumForcedGC])
	GCF bool
	// GC is the total execution cycles of GC ([runtime.MemStats.NumGC])
	GC bool
}

type prefix string

const (
	stats prefix = "stats"
	peak  prefix = "peak"
)

func param(alloc, total, sys uint64, gcf, gc uint32, prints ...Prints) params {
	var params params

	if prints[0].Alloc {
		params.allocs = &allocs{alloc: &alloc}
	}
	if prints[0].Total {
		params.allocs = &allocs{total: &total}
	}
	if prints[0].Sys {
		params.sys = &sys
	}
	if prints[0].GCF {
		params.gcs = &gcs{gcf: &gcf}
	}
	if prints[0].GC {
		params.gcs = &gcs{gc: &gc}
	}

	if prints[0].Allocs {
		params.allocs = &allocs{&alloc, &total}
	}
	if prints[0].Memory {
		params.allocs = &allocs{&alloc, &total}
		params.sys = &sys
	}
	if prints[0].GCs {
		params.gcs = &gcs{&gcf, &gc}
	}

	if prints[0].All {
		params.allocs = &allocs{&alloc, &total}
		params.sys = &sys
		params.gcs = &gcs{&gcf, &gc}
	}

	return params
}

func build(prefix prefix, params params) (string, []any) {
	prefix = "memory " + prefix + ":"
	var (
		message []string
		args    []any
	)
	mi := uint64(math.Pow(1024, 2))

	if params.allocs != nil {
		var (
			parts  []string
			values []any
		)
		if params.allocs.alloc != nil {
			parts = append(parts, "alloc %dMiB")
			values = append(values, *params.allocs.alloc/mi)
		}
		if params.allocs.total != nil {
			parts = append(parts, "total %dMiB")
			values = append(values, *params.allocs.total/mi)
		}
		message = append(message, strings.Join(parts, " - "))
		args = append(args, values...)
	}

	if params.sys != nil {
		message = append(message, "sys %dMiB")
		args = append(args, *params.sys/mi)
	}

	if params.gcs != nil {
		var (
			parts  []string
			values []any
		)
		if params.gcs.gcf != nil {
			parts = append(parts, "gcf %dc")
			values = append(values, *params.gcs.gcf)
		}
		if params.gcs.gc != nil {
			parts = append(parts, "gc %dc")
			values = append(values, *params.gcs.gc)
		}
		message = append(message, strings.Join(parts, " - "))
		args = append(args, values...)
	}

	return string(prefix) + " " + strings.Join(message, " / "), args
}

func print(prefix prefix, alloc, total, sys uint64, gcf, gc uint32, prints ...Prints) {
	if len(prints) > 0 {
		format, args := build(prefix, param(alloc, total, sys, gcf, gc, prints...))
		Printer(format+"\n", args...)
	}
}
