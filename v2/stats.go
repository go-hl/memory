package memory

import "runtime"

func mem(mem *runtime.MemStats) (uint64, uint64, uint64, uint32, uint32) {
	runtime.ReadMemStats(mem)
	return mem.Alloc, mem.TotalAlloc, mem.Sys, mem.NumForcedGC, mem.NumGC
}

// Stats can prints and returns the:
//  1. Atual memory allocated in Heap ([runtime.MemStats.Alloc])
//  1. Total memory allocated in Heap throughout the execution ([runtime.MemStats.TotalAlloc])
//  1. Reserved memory in OS for the programa ([runtime.MemStats.Sys])
//  1. Forced execution cycles of GC ([runtime.MemStats.NumForcedGC])
//  1. Total execution cycles of GC ([runtime.MemStats.NumGC])
func Stats(prints ...Prints) (alloc, total, sys uint64, gcf, gc uint32) {
	alloc, total, sys, gcf, gc = mem(&Mem)
	print(stats, alloc, total, sys, gcf, gc, prints...)
	return
}
