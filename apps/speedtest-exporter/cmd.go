package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	defaultPort             = 8080
	defaultCacheTime uint64 = 5
	defaultDuration         = time.Duration(defaultCacheTime * uint64(time.Minute))
	defaulLogLevel          = slog.LevelInfo
)

var (
	port           int
	addr           string
	cacheTime      uint64
	cacheDuration  time.Duration
	speedtestPath  string
	instance       string
	verboseLogging bool
)

// Initialize the needed flags for cli options
func init() {
	flag.IntVar(&port, "port", defaultPort, "Port for the webserver, default 8080")
	flag.Uint64Var(&cacheTime, "cacheTime", defaultCacheTime, "Time in minutes to cache speedtest output")
	flag.StringVar(&speedtestPath, "speedtest-path", "", "Specify speedtest executable to use, defaults to internal implementation")
	flag.StringVar(&instance, "instance", "", "Label added to all metrics for identification, defaults to hostname")
	flag.BoolVar(&verboseLogging, "v", false, "Enable verbose output")
}

// Parse Options not provided by the CLI Arguments from ENV
func parseEnv() {
	if port == defaultPort {
		if val, ok := os.LookupEnv("SPEEDTEST_PORT"); ok {
			var err error
			port, err = strconv.Atoi(val)
			if err != nil {
				slog.Error("Failed to convert SPEEDTEST_PORT to integer", "err", err)
				os.Exit(1)
			}
		}
	}
	if cacheTime == defaultCacheTime {
		if val, ok := os.LookupEnv("SPEEDTEST_CACHE_TIME"); ok {
			var err error
			cacheTime, err = strconv.ParseUint(val, 10, 0)
			if err != nil {
				slog.Error("Failed to convert SPEEDTEST_CACHE_TIME to integer", "err", err)
				os.Exit(1)
			}
		}
	}
	if speedtestPath == "" {
		if val, ok := os.LookupEnv("SPEEDTEST_PATH"); ok {
			speedtestPath = val
		}
	}
	if instance == "" {
		if val, ok := os.LookupEnv("SPEEDTEST_INSTANCE"); ok {
			instance = val
		}
	}
	if !verboseLogging {
		if val, ok := os.LookupEnv("SPEEDTEST_DEBUG"); ok {
			verboseLogging = strings.ToLower(val) == "true" || val == "1"
		}
	}
}

// Parse CLI Arguments and check the input.
// Provide dynamic default values where static ones are not possible
func parseFlags() {
	flag.Parse()

	parseEnv()

	if verboseLogging {
		logLevel.Set(slog.LevelDebug)
		slog.Info("Set log level to DEBUG")
	} else {
		logLevel.Set(defaulLogLevel)
	}

	addr = strings.Join([]string{":", strconv.Itoa(port)}, "")

	if instance == "" {
		slog.Info("No instance name provided, defaulting to hostname")
		hostname, err := os.Hostname()
		if err != nil {
			slog.Error("Failed to retrieve hostname, using localhost instead", "err", err)
			hostname = "localhost"
		}
		instance = hostname
	}

	var err error
	cacheDuration, err = time.ParseDuration(strconv.Itoa(int(cacheTime)) + "m")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unexcpected error parsing cache time: %v", err)
		os.Exit(1)
	}

	slog.Debug("The following settings are used",
		slog.Int("port", port),
		slog.Uint64("cache", cacheTime),
		slog.String("speedtestPath", speedtestPath),
		slog.String("instance", instance),
	)
}
