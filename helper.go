package stati

import (
	"runtime"
	"time"
)

func getStats() *statsResp {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)

	return &statsResp{
		HeapAlloc:     stats.HeapAlloc,
		HeapSys:       stats.HeapSys,
		HeapIdle:      stats.HeapIdle,
		HeapInuse:     stats.HeapInuse,
		HeapReleased:  stats.HeapReleased,
		HeapObjects:   stats.HeapObjects,
		Mallocs:       stats.Mallocs,
		Frees:         stats.Frees,
		NextGC:        stats.NextGC,
		LastGC:        time.Unix(0, int64(stats.LastGC)),
		GCSys:         stats.GCSys,
		Alloc:         stats.Alloc,
		NumGC:         stats.NumGC,
		NumForcedGC:   stats.NumForcedGC,
		GCCPUFraction: stats.GCCPUFraction,
		PauseTotalNs:  stats.PauseTotalNs,
		EnableGC:      stats.EnableGC,
		DebugGC:       stats.DebugGC,
		TotalAlloc:    stats.TotalAlloc,
		Sys:           stats.Sys,
		OtherSys:      stats.OtherSys,
		Lookups:       stats.Lookups,
		StackInuse:    stats.StackInuse,
		StackSys:      stats.StackSys,
		MSpanInuse:    stats.MSpanInuse,
		MSpanSys:      stats.MSpanSys,
		MCacheInuse:   stats.MCacheInuse,
		MCacheSys:     stats.MCacheSys,
		BuckHashSys:   stats.BuckHashSys,
		NumGoroutines: runtime.NumGoroutine(),
	}
}
