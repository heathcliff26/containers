package main

import (
	"flag"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

type expectedResult struct {
	webroot      string
	port         int
	withoutIndex bool
}

type testMatrixCmd struct {
	Name   string
	Args   []string
	Env    map[string]string
	result expectedResult
}

func TestCmd(t *testing.T) {
	testMatrix := []testMatrixCmd{
		{
			Name: "Default",
			Args: []string{"-webroot", "/foo/bar"},
			Env:  nil,
			result: expectedResult{
				webroot:      "/foo/bar",
				port:         defaultPort,
				withoutIndex: false,
			},
		},
		{
			Name: "Args",
			Args: []string{"-webroot", "/foo/bar", "-port", "1234", "-no-index"},
			Env:  nil,
			result: expectedResult{
				webroot:      "/foo/bar",
				port:         1234,
				withoutIndex: true,
			},
		},
		{
			Name: "Env",
			Args: nil,
			Env: map[string]string{
				"SFILESERVER_WEBROOT":  "/foo/bar/baz",
				"SFILESERVER_PORT":     "5678",
				"SFILESERVER_NO_INDEX": "tRue",
			},
			result: expectedResult{
				webroot:      "/foo/bar/baz",
				port:         5678,
				withoutIndex: true,
			},
		},
		{
			Name: "ArgsOverrideEnv",
			Args: []string{"-webroot", "/foo", "-port", "1234", "-no-index"},
			Env: map[string]string{
				"SFILESERVER_WEBROOT":  "/foo/bar/baz",
				"SFILESERVER_PORT":     "5678",
				"SFILESERVER_NO_INDEX": "false",
			},
			result: expectedResult{
				webroot:      "/foo",
				port:         1234,
				withoutIndex: true,
			},
		},
	}

	for _, tCase := range testMatrix {
		t.Run(tCase.Name, func(t *testing.T) {
			t.Cleanup(func() {
				webroot = ""
				port = defaultPort
				withoutIndex = false
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

			assert.Equal(tCase.result.webroot, webroot)
			assert.Equal(tCase.result.port, port)
			assert.Equal(tCase.result.withoutIndex, withoutIndex)
		})
	}
}

func TestCmdWebrootMissing(t *testing.T) {
	if os.Getenv("RUN_CRASH_TEST") == "1" {
		t.Setenv("SFILESERVER_WEBROOT", "")
		parseFlags()
		// Should not reach here, ensure exit with 0 if it does
		os.Exit(0)
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestCmdWebrootMissing")
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
		t.Setenv("SFILESERVER_PORT", "not a number")
		t.Setenv("SFILESERVER_WEBROOT", "/foo/bar")
		parseFlags()
		// Should not reach here, ensure exit with 0 if it does
		os.Exit(0)
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestCmdMalformedPortEnvVariable")
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
