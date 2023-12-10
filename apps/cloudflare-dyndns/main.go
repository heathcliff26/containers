package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/heathcliff26/containers/apps/cloudflare-dyndns/client"
	"github.com/heathcliff26/containers/apps/cloudflare-dyndns/config"
	"github.com/heathcliff26/containers/apps/cloudflare-dyndns/server"
)

var (
	clientMode bool
	configPath string
)

// Initialize the needed flags for cli options
func init() {
	flag.BoolVar(&clientMode, "client", false, "Run in client mode")
	flag.StringVar(&configPath, "config", "", "Path to config file, can be empty when running as server")
}

func main() {
	flag.Parse()
	var mode string
	if clientMode {
		mode = config.MODE_CLIENT
	} else {
		mode = config.MODE_SERVER
	}
	config, err := config.LoadConfig(configPath, mode)
	if err != nil {
		slog.Error("Could not load configuration", slog.String("path", configPath), slog.String("err", err.Error()))
		os.Exit(1)
	}

	if clientMode {
		c, err := client.NewCloudflareClient(config.Client.Secret, config.Client.Proxy)
		if err != nil {
			slog.Error("Could not create new client", "err", err)
			os.Exit(1)
		}
		c.SetDomains(config.Client.Domains)
		c.Run(config.Client.IntervalTime)
	} else {
		s := server.NewServer(config.Server)
		err := s.Run()
		if err != nil {
			slog.Error("Failed to start the server", "err", err)
			os.Exit(1)
		}
	}
}
