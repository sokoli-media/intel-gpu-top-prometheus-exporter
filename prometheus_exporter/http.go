package prometheus_exporter

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log/slog"
	"net/http"
)

func RunHTTPServer(logger *slog.Logger) {
	go exportMetrics(logger)

	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		logger.Error("failed to run http server", err)
		return
	}
}