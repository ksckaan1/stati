<script lang="ts">
    import {Line} from 'svelte-chartjs';

    import {
        Chart as ChartJS,
        Title,
        Tooltip,
        LineElement,
        LinearScale,
        PointElement,
        CategoryScale
    } from 'chart.js';
    import type {ChartValue, ChartTable} from "$lib/models";
    import {humanFileSize} from "$lib/utils";


    ChartJS.register(
        Title,
        Tooltip,
        LineElement,
        LinearScale,
        PointElement,
        CategoryScale
    );

    export let title: string = "Untitled"
    export let unit: string = ""
    export let values: ChartValue[] = []
    export let labels: string[] = []

    export let xName: string = "Time"
    export let yName: string = "Count"
    let isExpanded: boolean = false
    let group = values.map(v => v.name)
    let data = {}
    let table: ChartTable[]
    let isHumanReadable: boolean = unit === "byte"

    const updateTable = () => {
        let buff: ChartTable[] = []
        if (values) {
            values.forEach((v, i) => {
                buff = [...buff, {
                    name: v.name,
                    color: v.color,
                    first: table ? table[i].first : v.values[v.values.length - 1],
                    last: v.values ? v.values[v.values.length - 1] : 0,
                    max: v.values ? v.values[v.values.length - 1] : 0,
                    min: v.values ? v.values[v.values.length - 1] : 0,
                }]
            })

            if (!table) {
                table = buff
            } else {
                buff.forEach((_, i) => {
                    if (buff[i].max > table[i].max) table[i].max = buff[i].max
                    if (buff[i].min < table[i].min) table[i].min = buff[i].min
                    table[i].last = buff[i].last
                })

                table = table
            }
        }
    }

    $: {
        if (data) {
            updateTable()
        }
    }

    $: data = {
        labels: labels,
        datasets: values.length > 0 ? values.filter(v => group.includes(v.name)).map(v => {
            return {
                label: v.name,
                lineTension: 0,
                backgroundColor: v.color,
                borderColor: v.color,
                borderCapStyle: 'butt',
                borderDash: [],
                borderDashOffset: 0.0,
                borderJoinStyle: 'mitter',
                pointBorderColor: "transparent",
                pointBackgroundColor: "transparent",
                data: v.values,
            }
        }) : [],
    }

    let options = {}

    $: options = {
        scales: {
            y: {
                ticks: {
                    callback: function (value: number, index: number, ticks: number) {
                        return isHumanReadable ? humanFileSize(value) : value
                    },
                }
            }
        },
        interaction: {
            intersect: false,
            mode: 'index',
        },
        plugins: {
            tooltip: {
                position: "nearest",
                callbacks: {
                    labelColors: [{backgroundColor: "red"}],
                    label: (context) => {
                        let value = isHumanReadable ? humanFileSize(context.parsed.y) : context.parsed.y;
                        return `${context.dataset.label}: ${value}`
                    },
                }
            }
        },
        responsive: true,
        animation: false,
        maintainAspectRatio: false
    }

</script>

<div class="box" class:h-[500px]={!isExpanded}>
    <div class="title">
        <div class="flex items-center justify-between gap-3">
            <span>
            {title}
        </span>
            {#if unit}
            <span class="text-sm opacity-50">
                ({unit})
            </span>
            {/if}
        </div>
        <div class="flex gap-3 justify-between items-center">
            {#if unit === "byte"}
                <button on:click={()=> isHumanReadable =!isHumanReadable}
                        class="ml-auto float-right text-sm bg-black/10 px-2 py-1 rounded border border-white/20"
                        class:bg-green-900={isHumanReadable}
                >
                    Human Readable
                </button>
            {/if}
            <button on:click={()=> isExpanded =!isExpanded}
                    class="ml-auto float-right text-sm bg-black/10 px-2 py-1 rounded border border-white/20">
                {#if isExpanded}less{:else}more{/if}
            </button>
        </div>
    </div>
    <div class="flex gap-3 items-center justify-center mt-2">
        {#each values as v}
            <label for={v.name} class="flex gap-1 items-center cursor-pointer">
                <input class="hidden" bind:group type="checkbox" value={v.name} id={v.name}>
                <div class="h-4 w-4" style={`background-color: ${group.includes(v.name) ? v.color: "gray"}`}>

                </div>
                <span class:line-through={!group.includes(v.name)}>
                    {v.name}
                </span>
            </label>
        {/each}
    </div>
    <div class="flex">
        <div class="flex justify-center items-center">
            <span class="block origin-left-top transform -rotate-90 text-sm">
                {yName}
            </span>
        </div>
        <div class="graph">
            <Line {data} {options}/>
        </div>
    </div>
    <div class="text-center w-full text-sm mb-5">
        {xName}
    </div>
    <div class="grid grid-cols-5 bg-white/5 border-y border-white/20 font-bold">
        <div class="px-3 py-2">label</div>
        <div class="px-3 py-2">min</div>
        <div class="px-3 py-2">max</div>
        <div class="px-3 py-2">first</div>
        <div class="px-3 py-2">last</div>
    </div>
    <div class="tablo">
        {#each table as v}
            <div class="grid grid-cols-5">
                <div class="flex gap-2 items-center px-3 py-2">
                    <div class="h-4 w-4" style={`background-color: ${v.color}`}>

                    </div>
                    <span>
                    {v.name}
                </span>
                </div>
                <div class="px-3 py-2">{isHumanReadable ? humanFileSize(v.min) : v.min}</div>
                <div class="px-3 py-2">{isHumanReadable ? humanFileSize(v.max) : v.max}</div>
                <div class="px-3 py-2">{isHumanReadable ? humanFileSize(v.first) : v.first}</div>
                <div class="px-3 py-2">{isHumanReadable ? humanFileSize(v.last) : v.last}</div>
            </div>
        {/each}
    </div>
</div>

<style lang="postcss">
    .box {
        @apply rounded bg-white/5 overflow-hidden border border-white/20;
    }

    .graph {
        @apply h-96 p-3 flex-1;
    }

    .title {
        @apply flex justify-between items-center text-alto px-3 py-2 text-xl bg-white/5 border-b border-white/20;
    }

    .tablo > *:nth-child(2n) {
        @apply bg-black/20;
    }
</style>