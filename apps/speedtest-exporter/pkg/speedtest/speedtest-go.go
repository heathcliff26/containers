package speedtest

import (
	"log/slog"

	"github.com/heathcliff26/containers/apps/speedtest-exporter/pkg/collector"
	"github.com/showwin/speedtest-go/speedtest"
)

type Speedtest struct {
}

// Create instance of Speedtest
func NewSpeedtest() *Speedtest {
	return &Speedtest{}
}

// Use the speedtest-go api to run a speedtest and parse the result
func (s *Speedtest) Speedtest() *collector.SpeedtestResult {
	var client = speedtest.New()

	serverList, err := client.FetchServers()
	if err != nil {
		slog.Error("Could not fetch server list", "error", err)
		return collector.NewFailedSpeedtestResult()
	}
	targets, err := serverList.FindServer([]int{})
	if err != nil {
		slog.Error("Failed to find closest server", "error", err)
		return collector.NewFailedSpeedtestResult()
	}
	if len(targets) != 1 {
		slog.Error("FindServer returned more than one server")
		return collector.NewFailedSpeedtestResult()
	}
	server := targets[0]

	err = server.TestAll()
	if err != nil {
		slog.Error("Failed to run speedtest", "error", err)
		return collector.NewFailedSpeedtestResult()
	}
	user, err := client.FetchUserInfo()
	if err != nil {
		slog.Error("Failed to fetch client information", "error", err)
		return collector.NewFailedSpeedtestResult()
	}

	dataUsed := convertBytesToMB(server.Context.GetTotalDownload()) + convertBytesToMB(server.Context.GetTotalUpload())

	slog.Info("Successfully ran speedtest", slog.Group("result"),
		slog.Int64("jitterLatency", server.Jitter.Milliseconds()),
		slog.Int64("ping", server.Latency.Milliseconds()),
		slog.Float64("downloadSpeed", server.DLSpeed),
		slog.Float64("uploadSpeed", server.ULSpeed),
		slog.Float64("dataUsed", dataUsed),
		slog.String("serverId", server.ID),
		slog.String("serverHost", server.Host),
		slog.String("isp", user.Isp),
		slog.String("IP", user.IP),
	)

	return collector.NewSpeedtestResult(float64(server.Jitter.Milliseconds()), float64(server.Latency.Milliseconds()), server.DLSpeed, server.ULSpeed, dataUsed, user.Isp, user.IP)
}
