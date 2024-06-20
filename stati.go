package stati

import (
	"embed"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ksckaan1/stati/templates"
)

//go:embed assets/*
var assets embed.FS

type Stati struct {
	title       string
	interval    time.Duration
	chartBuffer int
	addr        string
}

func New() *Stati {
	return &Stati{
		title:       "Realtime Stats",
		interval:    time.Second,
		chartBuffer: 100,
		addr:        ":3000",
	}
}

func (s *Stati) WithTitle(title string) *Stati {
	s.title = title
	return s
}

func (s *Stati) WithInterval(interval time.Duration) *Stati {
	s.interval = interval
	return s
}

func (s *Stati) WithChartBuffer(buffer int) *Stati {
	s.chartBuffer = buffer
	return s
}

func (s *Stati) WithAddr(addr string) *Stati {
	s.addr = addr
	return s
}

func (s *Stati) Start() error {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.FileServer(http.FS(assets)))
	mux.HandleFunc("GET /stati", s.StatsPage)
	mux.HandleFunc("GET /debug/stati/info", s.getStatiInfo)

	err := http.ListenAndServe(s.addr, mux)
	if err != nil {
		return fmt.Errorf("http: listen and serve : %w", err)
	}

	return nil
}

func (s *Stati) StatsPage(w http.ResponseWriter, r *http.Request) {
	page := templates.StatsPage(templates.StatsConfig{
		Title:       s.title,
		Interval:    s.interval.Milliseconds(),
		ChartBuffer: s.chartBuffer,
	})
	page.Render(r.Context(), w)
}

func (s *Stati) getStatiInfo(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	stats := getStats()

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(stats); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
