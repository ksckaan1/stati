package stati

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"time"
)

//go:embed frontend/build/*
var feDir embed.FS

type Stati struct {
	cfg *Config
}

func New(cfg ...*Config) *Stati {
	var c *Config

	if len(cfg) > 0 {
		c = cfg[0]
	} else {
		c = &Config{
			Name:          "Realtime Stats",
			FetchInterval: time.Second,
		}
	}

	return &Stati{cfg: c}
}

func (s *Stati) Start(port int) error {
	var err error
	mux := http.NewServeMux()
	target, err := fs.Sub(feDir, "frontend/build")
	if err != nil {
		return fmt.Errorf("stati / start : %w", err)
	}

	mux.Handle("/", http.FileServer(http.FS(target)))
	mux.HandleFunc("/debug/stati/config", s.getConfig)
	mux.HandleFunc("/debug/stati/info", s.getStatiInfo)

	if err = http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		return fmt.Errorf("stati / start : %w", err)
	}

	return nil
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

func (s *Stati) getConfig(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]any{
		"name":           s.cfg.Name,
		"fetch_interval": s.cfg.FetchInterval.Milliseconds(),
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
