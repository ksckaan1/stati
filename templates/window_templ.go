// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

type WindowConfig struct {
	Title    string
	Subtitle string
	ChartID  string
	Fields   []string
}

func ChartWindow(cfg WindowConfig) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-data=\"{ is_colapsed : true }\" class=\"w-full rounded bg-gray-900 border border-white/20 overflow-hidden\"><div class=\"w-full bg-gray-800 border-b border-white/20 p-2 flex items-center justify-between\"><div class=\"flex gap-x-2 items-center\"><h2>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(cfg.Title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/window.templ`, Line: 14, Col: 19}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if cfg.Subtitle != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"text-sm font-light opacity-50\">(")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(cfg.Subtitle)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/window.templ`, Line: 17, Col: 21}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(")</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"flex gap-x-3\"><button @click=\"is_colapsed = !is_colapsed\" class=\"cursor-pointer px-3 py-0.5 border border-white/20 rounded hover:opacity-70\" x-text=\"is_colapsed ? &#39;Show more&#39;: &#39;Show less&#39;\">Show more</button></div></div><div class=\"flex flex-col w-full\"><div class=\"h-64 w-full p-3\"><canvas class=\"w-full h-full\" id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(cfg.ChartID)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/window.templ`, Line: 33, Col: 50}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></canvas></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(cfg.Fields) > 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-show=\"!is_colapsed\" class=\"w-full stats [&amp;&gt;.row:nth-child(2n-1)]:bg-white/5 [&amp;&gt;.row]:border-b [&amp;&gt;.row]:border-white/20 [&amp;&gt;.row:last-child]:border-b-0 [&amp;&gt;.row:last-child]:pb-3\"><div class=\"grid grid-cols-5 bg-white/5 gap-6 border-y border-t-white/20 border-b-white/50 text-white/50 font-light pt-3 [&amp;&gt;div:first-child]:pl-4 [&amp;&gt;div:last-child]:pr-4\"><div>Field</div><div>Min</div><div>Avg</div><div>Max</div><div>Latest</div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, field := range cfg.Fields {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"grid row grid-cols-5 gap-6 border-y border-white/20 [&amp;&gt;div:first-child]:pl-4 [&amp;&gt;div:last-child]:pr-4\"><div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(field)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/window.templ`, Line: 60, Col: 15}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				if cfg.Subtitle == "byte" {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-data=\"{v: &#39;0&#39;}\" x-text=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var6 string
					templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs("v")
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/window.templ`, Line: 63, Col: 43}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" x-init=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var7 string
					templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs("$watch('$store.table." + field + "?.min, is_human_readable', value => { v = is_human_readable ? humanReadable(value): value})")
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/window.templ`, Line: 63, Col: 182}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">0</div><div x-data=\"{v: &#39;0&#39;}\" x-text=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var8 string
					templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs("v")
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/window.templ`, Line: 66, Col: 43}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" x-init=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var9 string
					templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs("$watch('$store.table." + field + "?.avg, is_human_readable', value => { v = is_human_readable ? humanReadable(value): value})")
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/window.templ`, Line: 66, Col: 182}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">0</div><div x-data=\"{v: &#39;0&#39;}\" x-text=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var10 string
					templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs("v")
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/window.templ`, Line: 69, Col: 43}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" x-init=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var11 string
					templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs("$watch('$store.table." + field + "?.max, is_human_readable', value => { v = is_human_readable ? humanReadable(value): value})")
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/window.templ`, Line: 69, Col: 182}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">0</div><div x-data=\"{v: &#39;0&#39;}\" x-text=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var12 string
					templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs("v")
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/window.templ`, Line: 72, Col: 43}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" x-init=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var13 string
					templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs("$watch('$store.table." + field + "?.latest, is_human_readable', value => { v = is_human_readable ? humanReadable(value): value})")
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/window.templ`, Line: 72, Col: 185}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">0</div>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-text=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var14 string
					templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs("$store.table." + field + "?.min")
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/window.templ`, Line: 76, Col: 55}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">0</div><div x-text=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var15 string
					templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs("$store.table." + field + "?.avg")
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/window.templ`, Line: 79, Col: 55}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">0</div><div x-text=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var16 string
					templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs("$store.table." + field + "?.max")
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/window.templ`, Line: 82, Col: 55}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">0</div><div x-text=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var17 string
					templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs("$store.table." + field + "?.latest")
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/window.templ`, Line: 85, Col: 58}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">0</div>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
