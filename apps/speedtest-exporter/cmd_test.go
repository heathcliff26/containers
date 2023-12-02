package main

import (
	"flag"
	"log/slog"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertLogLevel(t *testing.T, expectedLevel slog.Level) {
	if logLevel != nil {
		assert.Equal(t, expectedLevel, logLevel.Level())
	} else {
		t.Error("Variable logLevel is not initialized")
	}
}

func TestDefaultFlags(t *testing.T) {
	parseFlags()
	t.Cleanup(func() {
		port = defaultPort
		addr = ""
		cacheTime = defaultCacheTime
		speedtestPath = ""
		instance = ""
		verboseLogging = false
		logLevel.Set(defaulLogLevel)
	})

	defaultAddr := strings.Join([]string{":", strconv.Itoa(defaultPort)}, "")

	assert := assert.New(t)

	assert.Equal(defaultPort, port)
	assert.Equal(defaultAddr, addr)
	assert.Equal(defaultCacheTime, cacheTime)
	assert.Equal("", speedtestPath)
	assert.NotEqual("", instance)
	assert.Equal(false, verboseLogging)
	assertLogLevel(t, defaulLogLevel)
}

func TestFlags(t *testing.T) {
	args := []string{
		"-port", "1234",
		"-cacheTime", "56",
		"-speedtest-path", "/foo/bar/speedtest",
		"-instance", "foo",
		"-v",
	}
	err := flag.CommandLine.Parse(args)
	if err != nil {
		t.Fatalf("Failed to parse test args: %v", err)
	}
	t.Cleanup(func() {
		port = defaultPort
		addr = ""
		cacheTime = defaultCacheTime
		speedtestPath = ""
		instance = ""
		verboseLogging = false
		logLevel.Set(defaulLogLevel)
	})

	parseFlags()

	assert := assert.New(t)

	assert.Equal(1234, port)
	assert.Equal(":1234", addr)
	assert.Equal(56, cacheTime)
	assert.Equal("/foo/bar/speedtest", speedtestPath)
	assert.Equal("foo", instance)
	assert.Equal(true, verboseLogging)
	assertLogLevel(t, slog.LevelDebug)
}

func TestArgsFromEnv(t *testing.T) {
	t.Setenv("SPEEDTEST_PORT", "5678")
	t.Setenv("SPEEDTEST_CACHE_TIME", "90")
	t.Setenv("SPEEDTEST_PATH", "/foo/bar/baz/speedtest")
	t.Setenv("SPEEDTEST_INSTANCE", "foobar")
	t.Setenv("SPEEDTEST_DEBUG", "TrUe")

	parseFlags()
	t.Cleanup(func() {
		port = defaultPort
		addr = ""
		cacheTime = defaultCacheTime
		speedtestPath = ""
		instance = ""
		verboseLogging = false
		logLevel.Set(defaulLogLevel)
	})

	assert := assert.New(t)

	assert.Equal(5678, port)
	assert.Equal(":5678", addr)
	assert.Equal(90, cacheTime)
	assert.Equal("/foo/bar/baz/speedtest", speedtestPath)
	assert.Equal("foobar", instance)
	assert.Equal(true, verboseLogging)
	assertLogLevel(t, slog.LevelDebug)
}

func TestArgsOverwriteEnv(t *testing.T) {
	t.Setenv("SPEEDTEST_PORT", "5678")
	t.Setenv("SPEEDTEST_CACHE_TIME", "90")
	t.Setenv("SPEEDTEST_PATH", "/foo/bar/baz/speedtest")
	t.Setenv("SPEEDTEST_INSTANCE", "foobar")
	t.Setenv("SPEEDTEST_DEBUG", "TrUe")

	args := []string{
		"-port", "1234",
		"-cacheTime", "56",
		"-speedtest-path", "/foo/bar/speedtest",
		"-instance", "foo",
		"-v",
	}
	err := flag.CommandLine.Parse(args)
	if err != nil {
		t.Fatalf("Failed to parse test args: %v", err)
	}
	t.Cleanup(func() {
		port = defaultPort
		addr = ""
		cacheTime = defaultCacheTime
		speedtestPath = ""
		instance = ""
		verboseLogging = false
		logLevel.Set(defaulLogLevel)
	})

	parseFlags()

	assert := assert.New(t)

	assert.Equal(1234, port)
	assert.Equal(":1234", addr)
	assert.Equal(56, cacheTime)
	assert.Equal("/foo/bar/speedtest", speedtestPath)
	assert.Equal("foo", instance)
	assert.Equal(true, verboseLogging)
	assertLogLevel(t, slog.LevelDebug)
}
