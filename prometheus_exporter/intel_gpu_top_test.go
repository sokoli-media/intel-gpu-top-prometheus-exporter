package prometheus_exporter

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

type JSON map[string]any

func (j JSON) ToString() string {
	dumped, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}

	return string(dumped)
}

func TestParsingJson__ParsePeriod(t *testing.T) {
	payload := JSON{
		"period": JSON{
			"duration": 1033.398519,
			"unit":     "ms",
		},
	}
	metrics, err := metricsFromJson(payload.ToString())
	require.NoError(t, err)

	require.Equal(t, 1033.398519, metrics.Period.Duration)
	require.Equal(t, "ms", metrics.Period.Unit)
}

func TestParsingJson__ParseFrequency(t *testing.T) {
	payload := JSON{
		"frequency": JSON{
			"requested": 273.853692,
			"actual":    272.886011,
			"unit":      "MHz",
		},
	}

	metrics, err := metricsFromJson(payload.ToString())
	require.NoError(t, err)

	require.Equal(t, 273.853692, metrics.Frequency.Requested)
	require.Equal(t, 272.886011, metrics.Frequency.Actual)
	require.Equal(t, "MHz", metrics.Frequency.Unit)
}

func TestParsingJson__ParseInterrupts(t *testing.T) {
	payload := JSON{
		"interrupts": JSON{
			"count": 142.249091,
			"unit":  "irq/s",
		},
	}

	metrics, err := metricsFromJson(payload.ToString())
	require.NoError(t, err)

	require.Equal(t, 142.249091, metrics.Interrupts.Count)
	require.Equal(t, "irq/s", metrics.Interrupts.Unit)
}

func TestParsingJson__ParseRC6(t *testing.T) {
	payload := JSON{
		"rc6": JSON{
			"value": 0.123000,
			"unit":  "%",
		},
	}

	metrics, err := metricsFromJson(payload.ToString())
	require.NoError(t, err)

	require.Equal(t, 0.123000, metrics.RC6.Value)
	require.Equal(t, "%", metrics.RC6.Unit)
}

func TestParsingJson__ParsePower(t *testing.T) {
	payload := JSON{
		"power": JSON{
			"GPU":     1.182905,
			"Package": 24.891973,
			"unit":    "W",
		},
	}

	metrics, err := metricsFromJson(payload.ToString())
	require.NoError(t, err)

	require.Equal(t, 1.182905, metrics.Power.GPU)
	require.Equal(t, 24.891973, metrics.Power.Package)
	require.Equal(t, "W", metrics.Power.Unit)
}

func TestParsingJson__ParseIMCBandwidth(t *testing.T) {
	payload := JSON{
		"imc-bandwidth": JSON{
			"reads":  7443.950521,
			"writes": 3175.988280,
			"unit":   "MiB/s",
		},
	}

	metrics, err := metricsFromJson(payload.ToString())
	require.NoError(t, err)

	require.Equal(t, 7443.950521, metrics.IMCBandwidth.Reads)
	require.Equal(t, 3175.988280, metrics.IMCBandwidth.Writes)
	require.Equal(t, "MiB/s", metrics.IMCBandwidth.Unit)
}

func TestParsingJson__ParseEngines(t *testing.T) {
	payload := JSON{
		"engines": JSON{
			"Render/3D/0": JSON{
				"busy": 13.620188,
				"sema": 0.100000,
				"wait": 0.200000,
				"unit": "%",
			},
			"Blitter/0": JSON{
				"busy": 0.300000,
				"sema": 0.400000,
				"wait": 0.500000,
				"unit": "%",
			},
			"Video/0": JSON{
				"busy": 19.697508,
				"sema": 10.643619,
				"wait": 0.600000,
				"unit": "%",
			},
			"VideoEnhance/0": JSON{
				"busy": 0.700000,
				"sema": 0.800000,
				"wait": 0.900000,
				"unit": "%",
			},
		},
	}

	metrics, err := metricsFromJson(payload.ToString())
	require.NoError(t, err)

	require.Equal(t, 13.620188, metrics.Engines.Render3D.Busy)
	require.Equal(t, 0.100000, metrics.Engines.Render3D.Sema)
	require.Equal(t, 0.200000, metrics.Engines.Render3D.Wait)
	require.Equal(t, "%", metrics.Engines.Render3D.Unit)

	require.Equal(t, 0.300000, metrics.Engines.Blitter.Busy)
	require.Equal(t, 0.400000, metrics.Engines.Blitter.Sema)
	require.Equal(t, 0.500000, metrics.Engines.Blitter.Wait)
	require.Equal(t, "%", metrics.Engines.Blitter.Unit)

	require.Equal(t, 19.697508, metrics.Engines.Video.Busy)
	require.Equal(t, 10.643619, metrics.Engines.Video.Sema)
	require.Equal(t, 0.600000, metrics.Engines.Video.Wait)
	require.Equal(t, "%", metrics.Engines.Video.Unit)

	require.Equal(t, 0.700000, metrics.Engines.VideoEnhance.Busy)
	require.Equal(t, 0.800000, metrics.Engines.VideoEnhance.Sema)
	require.Equal(t, 0.900000, metrics.Engines.VideoEnhance.Wait)
	require.Equal(t, "%", metrics.Engines.VideoEnhance.Unit)
}
