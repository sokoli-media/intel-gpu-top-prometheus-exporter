package prometheus_exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"log/slog"
	"time"
)

var unitLabels = []string{"unit"}
var period = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "intel_gpu_top_period"}, unitLabels)

var frequencyRequested = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "intel_gpu_top_frequency_requested"}, unitLabels)
var frequencyActual = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "intel_gpu_top_frequency_actual"}, unitLabels)

var interrupts = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "intel_gpu_top_interrupts"}, unitLabels)
var rc6 = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "intel_gpu_top_rc6"}, unitLabels)
var powerGpu = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "intel_gpu_top_power_gpu"}, unitLabels)
var powerPackage = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "intel_gpu_top_power_package"}, unitLabels)
var imcBandwidthReads = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "intel_gpu_top_imc_bandwidth_reads"}, unitLabels)
var imcBandwidthWrites = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "intel_gpu_top_imc_bandwidth_writes"}, unitLabels)

var enginesRender3dBusy = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "intel_gpu_top_engines_render3d_busy"}, unitLabels)
var enginesRender3dSema = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "intel_gpu_top_engines_render3d_sema"}, unitLabels)
var enginesRender3dWait = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "intel_gpu_top_engines_render3d_wait"}, unitLabels)

var enginesBlitterBusy = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "intel_gpu_top_engines_blitter_busy"}, unitLabels)
var enginesBlitterSema = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "intel_gpu_top_engines_blitter_sema"}, unitLabels)
var enginesBlitterWait = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "intel_gpu_top_engines_blitter_wait"}, unitLabels)

var enginesVideoBusy = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "intel_gpu_top_engines_video_busy"}, unitLabels)
var enginesVideoSema = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "intel_gpu_top_engines_video_sema"}, unitLabels)
var enginesVideoWait = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "intel_gpu_top_engines_video_wait"}, unitLabels)

var enginesVideoEnhanceBusy = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "intel_gpu_top_engines_video_enhance_busy"}, unitLabels)
var enginesVideoEnhanceSema = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "intel_gpu_top_engines_video_enhance_sema"}, unitLabels)
var enginesVideoEnhanceWait = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "intel_gpu_top_engines_video_enhance_wait"}, unitLabels)

var updatedAt = promauto.NewGauge(prometheus.GaugeOpts{Name: "intel_gpu_top_updated_at"})
var errors = promauto.NewCounter(prometheus.CounterOpts{Name: "intel_gpu_top_errors"})

func exportMetrics(logger *slog.Logger) {
	interval := 5 * time.Second
	metricsChannel := make(chan IntelGpuMetrics)
	errorsChannel := make(chan error)
	go loadMetrics(logger, metricsChannel, errorsChannel, interval)

	logger.Info("waiting for metrics")
	for {
		select {
		case metric := <-metricsChannel:
			logger.Info("received metrics, publishing")

			gpuMetricsToPrometheusMetrics(metric)
		case <-errorsChannel:
			errors.Inc()
		}
	}
}

func gpuMetricsToPrometheusMetrics(metric IntelGpuMetrics) {
	period.WithLabelValues(metric.Period.Unit).Set(metric.Period.Duration)

	frequencyRequested.WithLabelValues(metric.Frequency.Unit).Set(metric.Frequency.Requested)
	frequencyActual.WithLabelValues(metric.Frequency.Unit).Set(metric.Frequency.Actual)

	interrupts.WithLabelValues(metric.Interrupts.Unit).Set(metric.Interrupts.Count)

	rc6.WithLabelValues(metric.RC6.Unit).Set(metric.RC6.Value)

	powerGpu.WithLabelValues(metric.Power.Unit).Set(metric.Power.GPU)
	powerPackage.WithLabelValues(metric.Power.Unit).Set(metric.Power.Package)

	imcBandwidthReads.WithLabelValues(metric.IMCBandwidth.Unit).Set(metric.IMCBandwidth.Reads)
	imcBandwidthWrites.WithLabelValues(metric.IMCBandwidth.Unit).Set(metric.IMCBandwidth.Writes)

	enginesRender3dBusy.WithLabelValues(metric.Engines.Render3D.Unit).Set(metric.Engines.Render3D.Busy)
	enginesRender3dSema.WithLabelValues(metric.Engines.Render3D.Unit).Set(metric.Engines.Render3D.Sema)
	enginesRender3dWait.WithLabelValues(metric.Engines.Render3D.Unit).Set(metric.Engines.Render3D.Wait)

	enginesBlitterBusy.WithLabelValues(metric.Engines.Blitter.Unit).Set(metric.Engines.Blitter.Busy)
	enginesBlitterSema.WithLabelValues(metric.Engines.Blitter.Unit).Set(metric.Engines.Blitter.Sema)
	enginesBlitterWait.WithLabelValues(metric.Engines.Blitter.Unit).Set(metric.Engines.Blitter.Wait)

	enginesVideoBusy.WithLabelValues(metric.Engines.Video.Unit).Set(metric.Engines.Video.Busy)
	enginesVideoSema.WithLabelValues(metric.Engines.Video.Unit).Set(metric.Engines.Video.Sema)
	enginesVideoWait.WithLabelValues(metric.Engines.Video.Unit).Set(metric.Engines.Video.Wait)

	enginesVideoEnhanceBusy.WithLabelValues(metric.Engines.VideoEnhance.Unit).Set(metric.Engines.VideoEnhance.Busy)
	enginesVideoEnhanceSema.WithLabelValues(metric.Engines.VideoEnhance.Unit).Set(metric.Engines.VideoEnhance.Sema)
	enginesVideoEnhanceWait.WithLabelValues(metric.Engines.VideoEnhance.Unit).Set(metric.Engines.VideoEnhance.Wait)

	updatedAt.SetToCurrentTime()
}
