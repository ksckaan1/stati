setTimeout(() => {
  document.querySelector("#splash").classList.add("hidden")
}, 1500)

document.addEventListener("alpine:init", () => {
  Alpine.store("table", {
    goroutine_num: {},
    sys: {},
    other_sys: {},
    heap_sys: {},
    heap_inuse: {},
    heap_alloc: {},
    heap_released: {},
    heap_idle: {},
    heap_objects: {},
    mallocs: {},
    frees: {},
    gc_sys: {},
    gc_num: {},
    gc_num_forced: {},
    gc_cpu_fraction: 0,
    total_pause: 0,
    last_gc: "None",
    next_gc_byte_target: 0,
  })
})
const dataset = document.currentScript.dataset
const cfg = {
  Interval: +dataset.interval,
  ChartBuffer: +dataset.chartBuffer,
  IsGCOn: !!dataset.isGcOn
}



let isPaused = false
let showWarning = false

const pause = () => {
  isPaused = true
}

const resume = () => {
  isPaused = false
}

const warningToast = document.getElementById("warning")

const options = {
  responsive: true,
  maintainAspectRatio: false,
  interaction: {
    intersect: false,
    mode: "nearest"
  },
}

const newChart = (chartID, fields) => {
  const ch = new Chart(document.getElementById(chartID), {
    type: 'line',
    data: {
      labels: [],
      datasets: fields.map(fieldName => {
        return {
          label: fieldName,
          data: [],
          borderWidth: 1,
          pointBorderColor: "transparent",
          pointBackgroundColor: "transparent",
        }
      })
    },
    options
  });
  return ch
}

const addValue = (chart, time, field, value) => {
  const index = chart.data.datasets.findIndex(d => d.label === field)
  if (index === -1) return;

  if (chart.data.datasets[index].data.length >= cfg.ChartBuffer) {
    chart.data.datasets[index].data.shift()
  }

  chart.data.datasets[index].data.push(value)

  try {
    const sum = chart.data.datasets[index].data.reduce((a, b) => a + b, 0);
    const avg = (sum / chart.data.datasets[index].data.length) || 0;

    Alpine.store("table")[field] = {
      latest: Number(value),
      min: Math.min(...chart.data.datasets[index].data),
      max: Math.max(...chart.data.datasets[index].data),
      avg: avg
    }
  } catch (err) {
    console.error(err)
  }
  if (chart.data.labels.length != chart.data.datasets[index].data.length) {
    chart.data.labels.push(time)
  }
}

const updateGC = (gcCpuFraction, totalPause, lastGc, nextGcByteTarget) => {
  try {
    Alpine.store("table").gc_cpu_fraction = gcCpuFraction
    Alpine.store("table").total_pause = totalPause
    Alpine.store("table").last_gc = formatTime(lastGc)
    Alpine.store("table").next_gc_byte_target = nextGcByteTarget
  } catch (err) {
    console.error(err)
  }
}

const goroutineChart = newChart("goroutine-chart", [
  "goroutine_num"
])

const sysChart = newChart("sys-chart", [
  "sys",
  "other_sys"
])

const heapByteChart = newChart("heap-byte-chart", [
  "heap_sys",
  "heap_inuse",
  "heap_alloc",
  "heap_released",
  "heap_idle"
])

const heapCountChart = newChart("heap-count-chart", [
  "heap_objects",
  "mallocs",
  "frees"
])

let gcByteChart;
let gcNumChart;

if (cfg.IsGCOn) {
  gcByteChart = newChart("gc-byte-chart", [
    "gc_sys"
  ])

  gcNumChart = newChart("gc-num-chart", [
    "gc_num",
    "gc_num_forced"
  ])
}

setInterval(async () => {
  if (isPaused) return;

  try {
    const resp = await fetch("/debug/stati/info")
    const body = await resp.json()

    const time = new Date().toLocaleTimeString()

    // Goroutine (num)
    addValue(goroutineChart, time, "goroutine_num", body.num_goroutines)

    // Sys (byte)
    addValue(sysChart, time, "sys", body.sys)
    addValue(sysChart, time, "other_sys", body.other_sys)

    // Heap (byte)
    addValue(heapByteChart, time, "heap_sys", body.heap_sys)
    addValue(heapByteChart, time, "heap_inuse", body.heap_inuse)
    addValue(heapByteChart, time, "heap_alloc", body.heap_alloc)
    addValue(heapByteChart, time, "heap_released", body.heap_released)
    addValue(heapByteChart, time, "heap_idle", body.heap_idle)

    // Heap (count)
    addValue(heapCountChart, time, "heap_objects", body.heap_objects)
    addValue(heapCountChart, time, "mallocs", body.mallocs)
    addValue(heapCountChart, time, "frees", body.frees)

    if (cfg.IsGCOn) {
      // GC (byte)
      addValue(gcByteChart, time, "gc_sys", body.gc_sys)

      // GC (num)
      addValue(gcNumChart, time, "gc_num", body.num_gc)
      addValue(gcNumChart, time, "gc_num_forced", body.num_forced_gc)
    }

    goroutineChart.update()
    sysChart.update()
    heapByteChart.update()
    heapCountChart.update()

    if (cfg.IsGCOn) {
      gcByteChart.update()
      gcNumChart.update()
      updateGC(body.gc_cpu_fraction, body.pause_total_ns, body.last_gc, body.next_gc)
    }


    if (showWarning) {
      showWarning = false
      warningToast.classList.remove("block")
      warningToast.classList.add("hidden")
    }
  } catch (err) {
    console.error(err)
    if (!showWarning) {
      showWarning = true
      warningToast.classList.remove("hidden")
      warningToast.classList.add("block")
    }
  }
}, cfg.Interval);

const humanReadable = (bytes) => {
  const units = ["B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB", "YiB"]

  let unitIndex = 0
  while (bytes >= 1024) {
    bytes /= 1024
    unitIndex++
  }

  return `${bytes.toFixed(2)} ${units[unitIndex]}`
}

const formatTime = (time) => {
  if (time.includes("1970")) return "None";

  const ymd = time.split("T")[0]
  const hms = time.split("T")[1]


  return `${ymd}<br>${hms}`
}