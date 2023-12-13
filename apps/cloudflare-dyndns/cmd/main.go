package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/heathcliff26/containers/apps/cloudflare-dyndns/pkg/client"
	"github.com/heathcliff26/containers/apps/cloudflare-dyndns/pkg/config"
	"github.com/heathcliff26/containers/apps/cloudflare-dyndns/pkg/dyndns"
	"github.com/heathcliff26/containers/apps/cloudflare-dyndns/pkg/relay"
	"github.com/heathcliff26/containers/apps/cloudflare-dyndns/pkg/server"
)

var (
	mode       string
	configPath string
	env        bool
)

// Initialize the needed flags for cli options
func init() {
	flag.StringVar(&mode, "mode", config.MODE_SERVER, "Set what mode to run, options are \""+config.MODE_SERVER+"\", \""+config.MODE_CLIENT+"\" and \""+config.MODE_RELAY+"\"")
	flag.StringVar(&configPath, "config", "", "Path to config file, can be empty when running in mode "+config.MODE_SERVER)
	flag.BoolVar(&env, "env", false, "Used together with -config, when set will expand enviroment variables in config")
}

func main() {
	flag.Parse()

	cfg, err := config.LoadConfig(configPath, mode, env)
	if err != nil {
		slog.Error("Could not load configuration", slog.String("path", configPath), slog.String("err", err.Error()))
		os.Exit(1)
	}

	switch mode {
	case config.MODE_SERVER:
		s := server.NewServer(cfg.Server)
		err := s.Run()
		if err != nil {
			slog.Error("Failed to start the server", "err", err)
			os.Exit(1)
		}
	case config.MODE_CLIENT:
		c, err := client.NewCloudflareClient(cfg.Client.Token, cfg.Client.Proxy)
		if err != nil {
			slog.Error("Could not create new client", "err", err)
			os.Exit(1)
		}
		runClient(c, cfg.Client)
	case config.MODE_RELAY:
		r, err := relay.NewRelay(cfg.Client.Token, cfg.Client.Proxy, cfg.Client.Endpoint)
		if err != nil {
			slog.Error("Could not create new client", "err", err)
			os.Exit(1)
		}
		runClient(r, cfg.Client)
	default:
		slog.Error("Unknown mode", slog.String("mode", mode))
		os.Exit(1)
	}
}

func runClient(c dyndns.Client, cfg config.ClientConfig) {
	c.Data().SetDomains(cfg.Domains)
	dyndns.Run(c, cfg.IntervalTime)
}
