interface IStatInfoResp {
    sys: number // inuse
    mallocs: number // inuse
    frees: number // inuse
    heap_alloc: number // inuse
    heap_sys: number // inuse
    heap_idle: number // inuse
    heap_inuse: number // inuse
    heap_released: number // inuse
    heap_objects: number // inuse
    gc_sys: number // inuse
    other_sys: number // inuse
    next_gc: number // inuse
    last_gc: number // inuse
    pause_total_ns: number // inuse
    num_gc: number // inuse
    num_forced_gc: number // inuse
    gc_cpu_fraction: number // inuse
    enable_gc: boolean // inuse
    num_goroutines: number // inuse

    debug_gc: boolean
    stack_inuse: number
    stack_sys: number
    m_span_inuse: number
    m_span_sys: number
    m_cache_inuse: number
    m_cache_sys: number
    buck_hash_sys: number
    lookups: number
    alloc: number
    total_alloc: number
}

interface IStatConfigResp {
    name: string
    fetch_interval: number
}

type ChartValue = {
    name: string
    color: string
    values: number[]
}

type ChartTable = {
    name: string
    color: string
    first: number
    last: number
    max: number
    min: number
}

export type {IStatInfoResp, IStatConfigResp, ChartValue, ChartTable}