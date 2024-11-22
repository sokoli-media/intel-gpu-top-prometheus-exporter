package prometheus_exporter

import (
	"bufio"
	"encoding/json"
	"log/slog"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type IntelGpuMetrics struct {
	Period struct {
		Duration float64 `json:"duration"`
		Unit     string  `json:"unit"`
	} `json:"period"`
	Frequency struct {
		Requested float64 `json:"requested"`
		Actual    float64 `json:"actual"`
		Unit      string  `json:"unit"`
	} `json:"frequency"`
	Interrupts struct {
		Count float64 `json:"count"`
		Unit  string  `json:"unit"`
	} `json:"interrupts"`
	RC6 struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	} `json:"rc6"`
	Power struct {
		GPU     float64 `json:"GPU"`
		Package float64 `json:"Package"`
		Unit    string  `json:"unit"`
	} `json:"power"`
	IMCBandwidth struct {
		Reads  float64 `json:"reads"`
		Writes float64 `json:"writes"`
		Unit   string  `json:"unit"`
	} `json:"imc-bandwidth"`
	Engines struct {
		Render3D struct {
			Busy float64 `json:"busy"`
			Sema float64 `json:"sema"`
			Wait float64 `json:"wait"`
			Unit string  `json:"unit"`
		} `json:"Render/3D/0"`
		Blitter struct {
			Busy float64 `json:"busy"`
			Sema float64 `json:"sema"`
			Wait float64 `json:"wait"`
			Unit string  `json:"unit"`
		} `json:"Blitter/0"`
		Video struct {
			Busy float64 `json:"busy"`
			Sema float64 `json:"sema"`
			Wait float64 `json:"wait"`
			Unit string  `json:"unit"`
		} `json:"Video/0"`
		VideoEnhance struct {
			Busy float64 `json:"busy"`
			Sema float64 `json:"sema"`
			Wait float64 `json:"wait"`
			Unit string  `json:"unit"`
		} `json:"VideoEnhance/0"`
	} `json:"engines"`
	Clients map[string]interface{} `json:"clients"`
}

func removeTabs(input string) string {
	re := regexp.MustCompile(`\t+`)
	return re.ReplaceAllString(input, " ")
}

func loadMetrics(logger *slog.Logger, metricsChannel chan IntelGpuMetrics, interval time.Duration) {
	intervalInMs := strconv.Itoa(int(interval.Milliseconds()))
	cmd := exec.Command("/usr/bin/intel_gpu_top", "-J", "-s", intervalInMs)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	logger.Info("starting the command")
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(stdout)
	var jsonBuilder strings.Builder

	logger.Info("waiting for the output with metrics")
	for scanner.Scan() {
		line := scanner.Text()
		if line == "{" {
			jsonBuilder.Reset()
			jsonBuilder.WriteString(line)
		} else if line == "}" {
			jsonBuilder.WriteString(line)

			jsonString := removeTabs(jsonBuilder.String())
			logger.Info("processing metrics", "metrics", jsonString)

			var metrics IntelGpuMetrics
			if err := json.Unmarshal([]byte(jsonString), &metrics); err != nil {
				panic(err)
			}

			metricsChannel <- metrics
		} else {
			jsonBuilder.WriteString(line)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
