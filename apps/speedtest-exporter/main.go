package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/heathcliff26/containers/apps/speedtest-exporter/collector"
	"github.com/heathcliff26/containers/apps/speedtest-exporter/speedtest"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var logLevel *slog.LevelVar

// Initialize the logger
func init() {
	logLevel = &slog.LevelVar{}
	opts := slog.HandlerOptions{
		Level: logLevel,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &opts))
	slog.SetDefault(logger)
}

// Handle requests to the webroot.
// Serves static, human-readable HTML that provides a link to /metrics
func ServerRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><body><h1>Welcome to speedtest-exporter</h1>Click <a href='/metrics'>here</a> to see metrics.</body></html>")
}

func createSpeedtest(path string) (collector.Speedtest, error) {
	if path == "" {
		return speedtest.NewSpeedtest(), nil
	} else {
		return speedtest.NewSpeedtestCLI(path)
	}
}

func main() {
	parseFlags()

	s, err := createSpeedtest(speedtestPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed initialize speedtest-cli: %v\n", err)
		os.Exit(1)
	}

	collector, err := collector.NewCollector(cacheDuration, instance, s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed create collector: %v\n", err)
		os.Exit(1)
	}

	reg := prometheus.NewRegistry()
	reg.MustRegister(collector)

	http.HandleFunc("/", ServerRootHandler)
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))

	slog.Info("Starting http server", slog.String("addr", addr))
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start http server: %v\n", err)
		os.Exit(1)
	}
}
