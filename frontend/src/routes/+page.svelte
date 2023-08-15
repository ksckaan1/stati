<script lang='ts'>
    import "../style/main.css"
    import GraphBox from "../components/GraphBox.svelte";
    import {getStatiConfig, getStatiInfo} from "$lib/requests";
    import {onMount} from "svelte";
    import {fly, fade} from "svelte/transition"
    import StatiLogo from "../components/StatiLogo.svelte";
    import type {ChartValue} from "$lib/models";

    let maxData = 100
    let pageState: "connecting" | "connected" | "disconnected" = "connecting"
    let showDisconnectedToast = false
    let showConnectedToast = false
    let isRunning = true
    let isGCEnabled = false

    let gc_cpu_fraction = 0
    let pause_total_ns = 0
    let last_gc: number;
    let next_gc_target = 0

    let name = "Realtime Stats Title"

    let labels: string[] = []
    let emptyValue: number[] = []
    for (let i = 0; i < maxData; i++) {
        labels.push("")
        emptyValue.push(0)
    }

    let heapByteValues: ChartValue[] = [
        {
            name: "heap_sys",
            color: "green",
            values: [...emptyValue]
        },
        {
            name: "heap_inuse",
            color: "pink",
            values: [...emptyValue]
        },
        {
            name: "heap_alloc",
            color: "skyblue",
            values: [...emptyValue]
        },
        {
            name: "heap_released",
            color: "coral",
            values: [...emptyValue]
        },
        {
            name: "heap_idle",
            color: "orange",
            values: [...emptyValue]
        }
    ]

    let heapCountValues: ChartValue[] = [
        {
            name: "heap_objects",
            color: "green",
            values: [...emptyValue]
        },
        {
            name: "mallocs",
            color: "pink",
            values: [...emptyValue]
        },
        {
            name: "frees",
            color: "skyblue",
            values: [...emptyValue]
        },
    ]

    let goroutineValues: ChartValue[] = [
        {
            name: "goroutine_num",
            color: "green",
            values: [...emptyValue]
        }
    ]

    let sysValues: ChartValue[] = [
        {
            name: "sys",
            color: "green",
            values: [...emptyValue]
        },
        {
            name: "other_sys",
            color: "pink",
            values: [...emptyValue]
        }
    ]

    let gcByteValues: ChartValue[] = [
        {
            name: "gc_sys",
            color: "green",
            values: [...emptyValue]
        }
    ]

    let gcNumValues: ChartValue[] = [
        {
            name: "gc_num",
            color: "green",
            values: [...emptyValue]
        },
        {
            name: "gc_num_forced",
            color: "pink",
            values: [...emptyValue]
        }
    ]

    onMount(() => {
        getStatiConfig().then(cfgResp => {
            name = cfgResp.name
            setInterval(() => {
                if (!isRunning) return
                getStatiInfo().then(resp => {
                    let time = new Date().toTimeString().split(" ")[0]
                    labels = [...labels.slice(-(maxData - 1)), time]

                    // Heap Byte
                    heapByteValues[0].values.push(resp.heap_sys)
                    heapByteValues[1].values.push(resp.heap_inuse)
                    heapByteValues[2].values.push(resp.heap_alloc)
                    heapByteValues[3].values.push(resp.heap_released)
                    heapByteValues[4].values.push(resp.heap_idle)
                    heapByteValues.forEach((_, i) => {
                        if (heapByteValues[i].values.length > maxData) {
                            heapByteValues[i].values = heapByteValues[i].values.slice(-maxData)
                        }
                    })
                    heapByteValues = heapByteValues

                    // Heap Count
                    heapCountValues[0].values.push(resp.heap_objects)
                    heapCountValues[1].values.push(resp.mallocs)
                    heapCountValues[2].values.push(resp.frees)
                    heapCountValues.forEach((_, i) => {
                        if (heapCountValues[i].values.length > maxData) {
                            heapCountValues[i].values = heapCountValues[i].values.slice(-maxData)
                        }
                    })
                    heapCountValues = heapCountValues

                    // Go routine
                    goroutineValues[0].values.push(resp.num_goroutines)
                    goroutineValues.forEach((_, i) => {
                        if (goroutineValues[i].values.length > maxData) {
                            goroutineValues[i].values = goroutineValues[i].values.slice(-maxData)
                        }
                    })
                    goroutineValues = goroutineValues

                    // Sys
                    sysValues[0].values.push(resp.sys)
                    sysValues[1].values.push(resp.other_sys)
                    sysValues.forEach((_, i) => {
                        if (sysValues[i].values.length > maxData) {
                            sysValues[i].values = sysValues[i].values.slice(-maxData)
                        }
                    })
                    sysValues = sysValues

                    // GC Byte
                    gcByteValues[0].values.push(resp.gc_sys)
                    gcByteValues.forEach((_, i) => {
                        if (gcByteValues[i].values.length > maxData) {
                            gcByteValues[i].values = gcByteValues[i].values.slice(-maxData)
                        }
                    })
                    gcByteValues = gcByteValues

                    // GC Num
                    gcNumValues[0].values.push(resp.num_gc)
                    gcNumValues[1].values.push(resp.num_forced_gc)
                    gcNumValues.forEach((_, i) => {
                        if (gcNumValues[i].values.length > maxData) {
                            gcNumValues[i].values = gcNumValues[i].values.slice(-maxData)
                        }
                    })
                    gcNumValues = gcNumValues

                    isGCEnabled = resp.enable_gc
                    gc_cpu_fraction = resp.gc_cpu_fraction
                    pause_total_ns = resp.pause_total_ns
                    last_gc = resp.last_gc
                    next_gc_target = resp.next_gc

                    if (pageState != "connected" || showDisconnectedToast == true) {
                        showConnectedToast = true
                        setTimeout(() => showConnectedToast = false, 3000)
                    }

                    pageState = "connected"
                    showDisconnectedToast = false
                }).catch(() => showDisconnectedToast = true)

            }, cfgResp.fetch_interval)
        }).catch(() => pageState = "disconnected")
    })

    const processLastGCTime = (t: string) => {
        if (t === "1970-01-01T02:00:00+02:00") {
            return "not yet run"
        }

        return t.split("T")[1]
    }

</script>

<svelte:head>
    <title>{name}</title>
</svelte:head>

{#if pageState === "connected"}
    <header class="mt-10 p-3 container mx-auto flex justify-between items-center" in:fade={{delay:200}}>
        <h1>
            {name}
        </h1>
        <button class="px-3 py-2 rounded border border-white/20 transition-all duration-200"
                class:bg-red-900={isRunning}
                class:bg-green-900={!isRunning} on:click={()=>isRunning = !isRunning}>
            {#if isRunning}Stop{:else}Run{/if}
        </button>
    </header>
    <section class="mb-10 grid grid-cols-1 gap-5 p-3 container mx-auto">
        <div in:fly={{y: 100, delay: 200}}>
            <GraphBox values={goroutineValues} labels={labels} title="Goroutine" unit="num" yName="Num"/>
        </div>
        <div in:fly={{y: 100, delay: 400}}>
            <GraphBox values={sysValues} labels={labels} title="Sys" unit="byte" yName="Bytes"/>
        </div>
        <div in:fly={{y: 100, delay: 600}}>
            <GraphBox values={heapByteValues} labels={labels} title="Heap" unit="byte" yName="Bytes"/>
        </div>
        <div in:fly={{y: 100, delay: 800}}>
            <GraphBox values={heapCountValues} labels={labels} title="Heap" unit="count" yName="Count"/>
        </div>
        <div in:fly={{y: 100, delay: 1000}}>
            <div class="rounded border border-white/20">
                <div class="p-3 flex justify-between items-center border-white/20" class:bg-green-900={isGCEnabled}
                     class:border-b={isGCEnabled}
                     class:bg-red-900={!isGCEnabled}>
                    <h1 class="text-lg font-bold">Garbage Collector</h1>
                    <div>
                        <span>
                            {#if isGCEnabled}Enabled{:else}Disabled{/if}
                        </span>
                    </div>
                </div>
                {#if isGCEnabled}
                    <div class="grid grid-cols-4 gap-5 p-3 bg-green-800">
                        <div class="rounded border border-white/20 p-3 bg-black/10">
                            <h2>GC CPU Fraction</h2>
                            <div class="text-lg font-bold">{gc_cpu_fraction}</div>
                        </div>
                        <div class="rounded border border-white/20 p-3 bg-black/10">
                            <h2>Total Pause</h2>
                            <div class="text-lg font-bold">{pause_total_ns}ns</div>
                        </div>
                        <div class="rounded border border-white/20 p-3 bg-black/10">
                            <h2>Last GC</h2>
                            <div class="text-lg font-bold">{processLastGCTime(last_gc.toString())}</div>
                        </div>
                        <div class="rounded border border-white/20 p-3 bg-black/10">
                            <h2>Next GC Byte Target</h2>
                            <div class="text-lg font-bold">{next_gc_target} bytes</div>
                        </div>
                    </div>
                    <div class="p-3 gap-5 grid grid-cols-1">
                        <GraphBox values={gcByteValues} labels={labels} title="GC" unit="byte" yName="Bytes"/>
                        <GraphBox values={gcNumValues} labels={labels} title="GC" unit="num" yName="Num"/>
                    </div>
                {/if}
            </div>
        </div>
    </section>
{:else if pageState === "connecting"}
    <div class="flex flex-col w-screen h-screen justify-center items-center gap-5" out:fade={{duration:200}}>
        <div class="w-64">
            <StatiLogo/>
        </div>
        <span>Connecting...</span>
    </div>
{:else}
    <div class="flex flex-col w-screen h-screen justify-center items-center gap-5">
        <div class="w-64">
            <StatiLogo/>
        </div>
        <span>could not connect</span>
        <button class="rounded border border-white/20 px-3 py-2" on:click={()=> location.reload()}>retry</button>
    </div>
{/if}

{#if showDisconnectedToast}
    <div transition:fly={{duration: 200, y : -50}}
         class="fixed bottom-3 right-3 flex gap-3 items-center justify-center bg-red-900 border border-white/20 py-2 px-4 rounded-lg">
        <span>Disconnected</span>
        <button class="rounded border border-white/20 px-3 py-2" on:click={()=> location.reload()}>retry</button>
    </div>
{/if}

{#if showConnectedToast}
    <div transition:fly={{duration: 200, y : -50}}
         class="fixed bottom-3 right-3 flex gap-3 items-center justify-center bg-green-900 border border-white/20 py-2 px-4 rounded-lg">
        <span>Connected</span>
    </div>
{/if}

<style lang="postcss">
    header h1 {
        @apply text-2xl font-bold;
    }
</style>
