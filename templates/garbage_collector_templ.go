// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func GarbageCollector(cfg StatsConfig) templ.Component {
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
		if cfg.IsGCOn {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-full border border-white/20 rounded overflow-hidden bg-green-900/20\"><div class=\"bg-green-800\"><div class=\"bg-green-900 px-3 py-1 border-b border-white/20\">Garbage Collector</div><div class=\"grid grid-cols-4 gap-5 p-5 border-b border-white/20\"><div class=\"bg-green-900 p-3 rounded border border-white/20\"><h2 class=\"text-white/50\">GC CPU Fraction</h2><div x-text=\"$store.table.gc_cpu_fraction\">0</div></div><div class=\"bg-green-900 p-3 rounded border border-white/20\"><h2 class=\"text-white/50\">Total Pause</h2><div x-text=\"$store.table.total_pause\">0ns</div></div><div class=\"bg-green-900 p-3 rounded border border-white/20\"><h2 class=\"text-white/50\">Last GC</h2><div x-html=\"$store.table.last_gc\">None</div></div><div class=\"bg-green-900 p-3 rounded border border-white/20\"><h2 class=\"text-white/50\">Next GC Byte Target</h2><div x-text=\"$store.table.next_gc_byte_target\">0</div></div></div></div><div class=\"p-5 flex flex-col gap-5\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-full border p-3 text-center border-white/20 rounded overflow-hidden bg-red-900/20\">Garbage Collector is off</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
