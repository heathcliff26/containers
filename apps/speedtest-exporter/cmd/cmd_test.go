package main

import (
	"flag"
	"log/slog"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type expectedResultCmd struct {
	port           int
	addr           string
	cacheTime      uint64
	cacheDuration  time.Duration
	speedtestPath  string
	instance       string
	verboseLogging bool
	logLevel       slog.Level
}

type testCasesCmd struct {
	Name   string
	Args   []string
	Env    map[string]string
	Result expectedResultCmd
}

func TestCmd(t *testing.T) {
	if logLevel == nil {
		t.Fatal("Variable logLevel is not initialized")
	}

	tMatrix := []testCasesCmd{
		{
			Name: "Default",
			Args: nil,
			Env:  nil,
			Result: expectedResultCmd{
				port:           defaultPort,
				addr:           strings.Join([]string{":", strconv.Itoa(defaultPort)}, ""),
				cacheTime:      defaultCacheTime,
				cacheDuration:  defaultDuration,
				speedtestPath:  "",
				instance:       "",
				verboseLogging: false,
				logLevel:       defaulLogLevel,
			},
		},
		{
			Name: "Args",
			Args: []string{
				"-port", "1234",
				"-cacheTime", "56",
				"-speedtest-path", "/foo/bar/speedtest",
				"-instance", "foo",
				"-v",
			},
			Env: nil,
			Result: expectedResultCmd{
				port:           1234,
				addr:           ":1234",
				cacheTime:      56,
				cacheDuration:  time.Duration(56 * time.Minute),
				speedtestPath:  "/foo/bar/speedtest",
				instance:       "foo",
				verboseLogging: true,
				logLevel:       slog.LevelDebug,
			},
		},
		{
			Name: "Env",
			Args: nil,
			Env: map[string]string{
				"SPEEDTEST_PORT":       "5678",
				"SPEEDTEST_CACHE_TIME": "90",
				"SPEEDTEST_PATH":       "/foo/bar/baz/speedtest",
				"SPEEDTEST_INSTANCE":   "foobar",
				"SPEEDTEST_DEBUG":      "TrUe",
			},
			Result: expectedResultCmd{
				port:           5678,
				addr:           ":5678",
				cacheTime:      90,
				cacheDuration:  time.Duration(90 * time.Minute),
				speedtestPath:  "/foo/bar/baz/speedtest",
				instance:       "foobar",
				verboseLogging: true,
				logLevel:       slog.LevelDebug,
			},
		},
		{
			Name: "ArgsOverwriteEnv",
			Args: []string{
				"-port", "1234",
				"-cacheTime", "56",
				"-speedtest-path", "/foo/bar/speedtest",
				"-instance", "foo",
				"-v",
			},
			Env: map[string]string{
				"SPEEDTEST_PORT":       "5678",
				"SPEEDTEST_CACHE_TIME": "90",
				"SPEEDTEST_PATH":       "/foo/bar/baz/speedtest",
				"SPEEDTEST_INSTANCE":   "foobar",
				"SPEEDTEST_DEBUG":      "false",
			},
			Result: expectedResultCmd{
				port:           1234,
				addr:           ":1234",
				cacheTime:      56,
				cacheDuration:  time.Duration(56 * time.Minute),
				speedtestPath:  "/foo/bar/speedtest",
				instance:       "foo",
				verboseLogging: true,
				logLevel:       slog.LevelDebug,
			},
		},
	}

	for _, tCase := range tMatrix {
		t.Run(tCase.Name, func(t *testing.T) {
			t.Cleanup(func() {
				port = defaultPort
				addr = ""
				cacheTime = defaultCacheTime
				cacheDuration = 0
				speedtestPath = ""
				instance = ""
				verboseLogging = false
				logLevel.Set(defaulLogLevel)
			})

			if tCase.Args != nil {
				err := flag.CommandLine.Parse(tCase.Args)
				if err != nil {
					t.Fatalf("Failed to parse test args: %v", err)
				}
			}
			for key, val := range tCase.Env {
				t.Setenv(key, val)
			}

			parseFlags()

			assert := assert.New(t)
			r := tCase.Result

			assert.Equal(r.port, port)
			assert.Equal(r.addr, addr)
			assert.Equal(r.cacheTime, cacheTime)
			assert.Equal(r.cacheDuration, cacheDuration)
			assert.Equal(r.speedtestPath, speedtestPath)
			if r.instance == "" {
				assert.NotEqual("", instance)
			} else {
				assert.Equal(r.instance, instance)
			}
			assert.Equal(r.verboseLogging, verboseLogging)
			assert.Equal(r.logLevel, logLevel.Level())
		})
	}
}

func execExitTest(t *testing.T, test string) {
	cmd := exec.Command(os.Args[0], "-test.run="+test)
	cmd.Env = append(os.Environ(), "RUN_CRASH_TEST=1")
	err := cmd.Run()
	if err == nil {
		t.Fatal("Process exited without error")
	}
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestCmdMalformedPortEnvVariable(t *testing.T) {
	if os.Getenv("RUN_CRASH_TEST") == "1" {
		t.Setenv("SPEEDTEST_PORT", "not a number")
		parseFlags()
		// Should not reach here, ensure exit with 0 if it does
		os.Exit(0)
	}
	execExitTest(t, "TestCmdMalformedPortEnvVariable")
}

func TestCmdMalformedCacheTimeEnvVariable(t *testing.T) {
	if os.Getenv("RUN_CRASH_TEST") == "1" {
		t.Setenv("SPEEDTEST_CACHE_TIME", "not a number")
		parseFlags()
		// Should not reach here, ensure exit with 0 if it does
		os.Exit(0)
	}
	execExitTest(t, "TestCmdMalformedCacheTimeEnvVariable")
}
