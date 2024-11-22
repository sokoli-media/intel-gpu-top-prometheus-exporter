package main

import (
	"intel-gpu-top-prometheus-exporter/prometheus_exporter"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	prometheus_exporter.RunHTTPServer(logger)
}
