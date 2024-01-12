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
	flag.StringVar(&sslCert, "cert", "", "SFILESERVER_CERT: SSL certificate to use, needs key as well. Default is no ssl.")
	flag.StringVar(&sslKey, "key", "", "SFILESERVER_KEY: SSL private key to use, needs cert as well. Default is no ssl.")
	flag.BoolVar(&withoutIndex, "no-index", false, "SFILESERVER_NO_INDEX: Do not serve an index for directories, return index.html or 404 instead")
	flag.BoolVar(&enableLogging, "log", false, "SFILESERVER_LOG: Enable logging requests")
}

func envBool(target *bool, name string) {
	if val, ok := os.LookupEnv(name); ok {
		*target = strings.ToLower(val) == "true" || val == "1"
	}
}

func envString(target *string, name string) {
	if val, ok := os.LookupEnv(name); ok {
		*target = val
	}
}

// Parse Options not provided by the CLI Arguments from ENV
func parseEnv() {
	if webroot == "" {
		envString(&webroot, "SFILESERVER_WEBROOT")
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

	if sslCert == "" {
		envString(&sslCert, "SFILESERVER_CERT")
	}

	if sslKey == "" {
		envString(&sslKey, "SFILESERVER_KEY")
	}

	if !withoutIndex {
		envBool(&withoutIndex, "SFILESERVER_NO_INDEX")
	}

	if !enableLogging {
		envBool(&enableLogging, "SFILESERVER_LOG")
	}
}

// Parse CLI Arguments and check the input.
func parseFlags() {
	flag.Parse()
	parseEnv()

	if webroot == "" {
		log.Fatal("No Webroot: Either -webroot or SFILESERVER_WEBROOT need to be set")
	}
	if (sslCert != "" && sslKey == "") || (sslCert == "" && sslKey != "") {
		log.Fatal("When using ssl need both -cert and -key to be set")
	}
	log.Printf("Settings: webroot=%s, port=%d, sslCert=%s, sslKey=%s, no-index=%t, enableLogging=%t", webroot, port, sslCert, sslKey, withoutIndex, enableLogging)
}
