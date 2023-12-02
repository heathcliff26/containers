package speedtest

import "testing"

func TestRunSpeedtestForGo(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	s := NewSpeedtest()
	result := s.Speedtest()
	if !result.Success() {
		t.Fatal("Speedtest returned with failure")
	}
}
