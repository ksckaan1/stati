package stati

import "time"

type Config struct {
	Name          string        `json:"name"`
	FetchInterval time.Duration `json:"fetch_interval"`
}

type statsResp struct {
	Alloc         uint64    `json:"alloc"`
	TotalAlloc    uint64    `json:"total_alloc"`
	Sys           uint64    `json:"sys"`
	Lookups       uint64    `json:"lookups"`
	Mallocs       uint64    `json:"mallocs"`
	Frees         uint64    `json:"frees"`
	HeapAlloc     uint64    `json:"heap_alloc"`
	HeapSys       uint64    `json:"heap_sys"`
	HeapIdle      uint64    `json:"heap_idle"`
	HeapInuse     uint64    `json:"heap_inuse"`
	HeapReleased  uint64    `json:"heap_released"`
	HeapObjects   uint64    `json:"heap_objects"`
	StackInuse    uint64    `json:"stack_inuse"`
	StackSys      uint64    `json:"stack_sys"`
	MSpanInuse    uint64    `json:"m_span_inuse"`
	MSpanSys      uint64    `json:"m_span_sys"`
	MCacheInuse   uint64    `json:"m_cache_inuse"`
	MCacheSys     uint64    `json:"m_cache_sys"`
	BuckHashSys   uint64    `json:"buck_hash_sys"`
	GCSys         uint64    `json:"gc_sys"`
	OtherSys      uint64    `json:"other_sys"`
	NextGC        uint64    `json:"next_gc"`
	LastGC        time.Time `json:"last_gc"`
	PauseTotalNs  uint64    `json:"pause_total_ns"`
	NumGC         uint32    `json:"num_gc"`
	NumForcedGC   uint32    `json:"num_forced_gc"`
	GCCPUFraction float64   `json:"gc_cpu_fraction"`
	EnableGC      bool      `json:"enable_gc"`
	DebugGC       bool      `json:"debug_gc"`
	NumGoroutines int       `json:"num_goroutines"`
}
