package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

const defaultPort = 8080

func init() {
	flag.StringVar(&webroot, "webroot", "", "SFILESERVER_WEBROOT: Required, root directory to serve files from")
	flag.IntVar(&port, "port", defaultPort, "SFILESERVER_PORT: Specify port for the fileserver to listen on")
	flag.BoolVar(&withoutIndex, "no-index", false, "SFILESERVER_NO_INDEX: Do not serve an index for directories, return index.html or 404 instead")
	flag.BoolVar(&debug, "debug", false, "SFILESERVER_DEBUG: Enable debug output")
}

func envBool(target *bool, name string) {
	if val, ok := os.LookupEnv(name); ok {
		*target = strings.ToLower(val) == "true" || val == "1"
	}
}

// Parse Options not provided by the CLI Arguments from ENV
func parseEnv() {
	if webroot == "" {
		if val, ok := os.LookupEnv("SFILESERVER_WEBROOT"); ok {
			webroot = val
		}
	}

	if port == defaultPort {
		if val, ok := os.LookupEnv("SFILESERVER_PORT"); ok {
			var err error
			port, err = strconv.Atoi(val)
			if err != nil {
				log.Fatalf("Could not parse SFILESERVER_PORT: %v", err)
			}
		}
	}

	if !withoutIndex {
		envBool(&withoutIndex, "SFILESERVER_NO_INDEX")
	}

	if !debug {
		envBool(&debug, "SFILESERVER_DEBUG")
	}
}

// Parse CLI Arguments and check the input.
func parseFlags() {
	flag.Parse()
	parseEnv()

	if webroot == "" {
		log.Fatal("No Webroot: Either -webroot or SFILESERVER_WEBROOT need to be set")
	}
	log.Printf("Settings: webroot=%s, port=%d, no-index=%t, debug=%t", webroot, port, withoutIndex, debug)
}