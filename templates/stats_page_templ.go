// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

type StatsConfig struct {
	Title       string
	ChartBuffer int
	Interval    int64
}

func StatsPage(cfg StatsConfig) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html><head><title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(cfg.Title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/stats_page.templ`, Line: 14, Col: 15}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><script src=\"/assets/tailwindcss.js\"></script></head><body class=\"bg-gray-900 text-zinc-300\"><main class=\"max-w-7xl mx-auto px-5 mt-10\"><header class=\"flex justify-between items-center\"><div class=\"text-2xl font-bold\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(cfg.Title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/stats_page.templ`, Line: 23, Col: 48}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><button id=\"pause-resume-btn\" class=\"bg-red-900 border border-white/20 rounded px-3 py-1\">Pause</button></header><div class=\"flex flex-col gap-5 my-5\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ChartWindow(WindowConfig{
			Title:    "Goroutine",
			Subtitle: "num",
			ChartID:  "goroutine-chart",
			Fields: []string{
				"goroutine_num",
			},
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ChartWindow(WindowConfig{
			Title:    "Sys",
			Subtitle: "byte",
			ChartID:  "sys-chart",
			Fields: []string{
				"sys",
				"other_sys",
			},
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ChartWindow(WindowConfig{
			Title:    "Heap",
			Subtitle: "byte",
			ChartID:  "heap-byte-chart",
			Fields: []string{
				"heap_sys",
				"heap_inuse",
				"heap_alloc",
				"heap_released",
				"heap_idle",
			},
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ChartWindow(WindowConfig{
			Title:    "Heap",
			Subtitle: "count",
			ChartID:  "heap-count-chart",
			Fields: []string{
				"heap_objects",
				"mallocs",
				"frees",
			},
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-full border border-white/20 rounded overflow-hidden\"><div class=\"bg-green-800\"><div class=\"bg-green-900 px-3 py-1 border-b border-white/20\">Garbage Collector</div><div class=\"grid grid-cols-4 gap-5 p-5 border-b border-white/20\"><div class=\"bg-green-900 p-3 rounded border border-white/20\"><h2>GC CPU Fraction</h2><div id=\"gc_cpu_fraction\">0</div></div><div class=\"bg-green-900 p-3 rounded border border-white/20\"><h2>Total Pause</h2><div id=\"total_pause\">0</div></div><div class=\"bg-green-900 p-3 rounded border border-white/20\"><h2>Last GC</h2><div id=\"last_gc\">None</div></div><div class=\"bg-green-900 p-3 rounded border border-white/20\"><h2>Next GC Byte Target</h2><div id=\"next_gc_byte_target\">0</div></div></div></div><div class=\"p-5 flex flex-col gap-5\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ChartWindow(WindowConfig{
			Title:    "GC",
			Subtitle: "byte",
			ChartID:  "gc-byte-chart",
			Fields: []string{
				"gc_sys",
			},
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ChartWindow(WindowConfig{
			Title:    "GC",
			Subtitle: "num",
			ChartID:  "gc-num-chart",
			Fields: []string{
				"gc_num",
				"gc_num_forced",
			},
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div></main><div id=\"warning\" class=\"fixed bottom-5 right-5 hidden p-5 bg-red-900 rounded border border-white/20 shadow-lg\">Error when fetching stats</div><script src=\"/assets/chart.js\"></script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = onLoad(cfg).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func onLoad(cfg StatsConfig) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_onLoad_8ab9`,
		Function: `function __templ_onLoad_8ab9(cfg){let isPaused = false
	let showWarning = false

	const warningToast = document.getElementById("warning")

	const pauseResumeBtn = document.getElementById("pause-resume-btn")
	pauseResumeBtn.addEventListener("click", (e) => {
		isPaused = !isPaused

		if (isPaused) {
			e.target.classList.remove("bg-red-900")
			e.target.classList.add("bg-green-900")
			e.target.innerText = "Resume"
		} else {
			e.target.classList.remove("bg-green-900")
			e.target.classList.add("bg-red-900")
			e.target.innerText = "Pause"
		}
	})

	const options = {
			responsive: true,
      maintainAspectRatio: false
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
			document.getElementById(field+"_latest").innerText = Number(value)
			document.getElementById(field+"_min").innerText = Math.min(...chart.data.datasets[index].data)
			document.getElementById(field+"_max").innerText = Math.max(...chart.data.datasets[index].data)

			const sum = chart.data.datasets[index].data.reduce((a, b) => a + b, 0);
			const avg = (sum / chart.data.datasets[index].data.length) || 0;
			document.getElementById(field+"_avg").innerText = avg

		} catch (err) {
			console.error(err)
		}
		if (chart.data.labels.length != chart.data.datasets[index].data.length) {
			chart.data.labels.push(time)
		}
	}

	const updateGC = (gcCpuFraction, totalPause, lastGc, nextGcByteTarget) => {
		try {
			document.getElementById("gc_cpu_fraction").innerText = gcCpuFraction
			document.getElementById("total_pause").innerText = totalPause
			document.getElementById("last_gc").innerText = lastGc.includes("1970") ? "None": lastGc
			document.getElementById("next_gc_byte_target").innerText = nextGcByteTarget
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

	const gcByteChart = newChart("gc-byte-chart", [
		"gc_sys"
		])
	
	const gcNumChart = newChart("gc-num-chart", [
		"gc_num",
		"gc_num_forced"
		])

	setInterval(async ()=>{
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

			// GC (byte)
			addValue(gcByteChart, time, "gc_sys", body.gc_sys)

			// GC (num)
			addValue(gcNumChart, time, "gc_num", body.num_gc)
			addValue(gcNumChart, time, "gc_num_forced", body.num_forced_gc)


			goroutineChart.update()
			sysChart.update()
			heapByteChart.update()
			heapCountChart.update()
			gcByteChart.update()
			gcNumChart.update()

			console.log(body)

			updateGC(body.gc_cpu_fraction, body.pause_total_ns,body.last_gc,body.next_gc)

			if (showWarning) {
				showWarning = false
				warningToast.classList.remove("block")
				warningToast.classList.add("hidden")
			}
		} catch (err) {
			if (!showWarning) {
				showWarning = true
				warningToast.classList.remove("hidden")
				warningToast.classList.add("block")
			}
		}
	},cfg.Interval);
}`,
		Call:       templ.SafeScript(`__templ_onLoad_8ab9`, cfg),
		CallInline: templ.SafeScriptInline(`__templ_onLoad_8ab9`, cfg),
	}
}