package main

import (
	"flag"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

type expectedResult struct {
	webroot         string
	port            int
	sslKey, sslCert string
	withoutIndex    bool
	debug           bool
}

func TestCmd(t *testing.T) {
	testMatrix := []struct {
		Name   string
		Args   []string
		Env    map[string]string
		result expectedResult
	}{
		{
			Name: "Default",
			Args: []string{"-webroot", "/foo/bar"},
			Env:  nil,
			result: expectedResult{
				webroot:      "/foo/bar",
				port:         defaultPort,
				withoutIndex: false,
				debug:        false,
			},
		},
		{
			Name: "Args",
			Args: []string{"-webroot", "/foo/bar", "-port", "1234", "-no-index", "-debug"},
			Env:  nil,
			result: expectedResult{
				webroot:      "/foo/bar",
				port:         1234,
				withoutIndex: true,
				debug:        true,
			},
		},
		{
			Name: "Env",
			Args: nil,
			Env: map[string]string{
				"SFILESERVER_WEBROOT":  "/foo/bar/baz",
				"SFILESERVER_PORT":     "5678",
				"SFILESERVER_NO_INDEX": "tRue",
				"SFILESERVER_DEBUG":    "trUe",
			},
			result: expectedResult{
				webroot:      "/foo/bar/baz",
				port:         5678,
				withoutIndex: true,
				debug:        true,
			},
		},
		{
			Name: "ArgsOverrideEnv",
			Args: []string{"-webroot", "/foo", "-port", "1234", "-no-index", "-debug"},
			Env: map[string]string{
				"SFILESERVER_WEBROOT":  "/foo/bar/baz",
				"SFILESERVER_PORT":     "5678",
				"SFILESERVER_NO_INDEX": "false",
				"SFILESERVER_DEBUG":    "false",
			},
			result: expectedResult{
				webroot:      "/foo",
				port:         1234,
				withoutIndex: true,
				debug:        true,
			},
		},
		{
			Name: "ArgsWithSSL",
			Args: []string{"-webroot", "/foo", "-cert", "ssl.crt", "-key", "ssl.key"},
			Env:  nil,
			result: expectedResult{
				webroot: "/foo",
				port:    defaultPort,
				sslCert: "ssl.crt",
				sslKey:  "ssl.key",
			},
		},
		{
			Name: "EnvWithSSL",
			Args: nil,
			Env: map[string]string{
				"SFILESERVER_WEBROOT": "/foo",
				"SFILESERVER_CERT":    "env.crt",
				"SFILESERVER_KEY":     "env.key",
			},
			result: expectedResult{
				webroot: "/foo",
				port:    defaultPort,
				sslCert: "env.crt",
				sslKey:  "env.key",
			},
		},
	}

	for _, tCase := range testMatrix {
		t.Run(tCase.Name, func(t *testing.T) {
			t.Cleanup(func() {
				webroot = ""
				port = defaultPort
				sslCert = ""
				sslKey = ""
				withoutIndex = false
				debug = false
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
			assert.Equal(tCase.result.sslCert, sslCert)
			assert.Equal(tCase.result.sslKey, sslKey)
			assert.Equal(tCase.result.withoutIndex, withoutIndex)
			assert.Equal(tCase.result.debug, debug)
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

func TestCmdWebrootMissing(t *testing.T) {
	if os.Getenv("RUN_CRASH_TEST") == "1" {
		t.Setenv("SFILESERVER_WEBROOT", "")
		parseFlags()
		// Should not reach here, ensure exit with 0 if it does
		os.Exit(0)
	}
	execExitTest(t, "TestCmdWebrootMissing")
}

func TestCmdMalformedPortEnvVariable(t *testing.T) {
	if os.Getenv("RUN_CRASH_TEST") == "1" {
		t.Setenv("SFILESERVER_PORT", "not a number")
		t.Setenv("SFILESERVER_WEBROOT", "/foo/bar")
		parseFlags()
		// Should not reach here, ensure exit with 0 if it does
		os.Exit(0)
	}
	execExitTest(t, "TestCmdMalformedPortEnvVariable")
}

func TestCmdIncompleteSSL(t *testing.T) {
	t.Run("NoCert", func(t *testing.T) {
		if os.Getenv("RUN_CRASH_TEST") == "1" {
			t.Setenv("SFILESERVER_WEBROOT", "/foo/bar")
			t.Setenv("SFILESERVER_KEY", "ssl.key")
			parseFlags()
			// Should not reach here, ensure exit with 0 if it does
			os.Exit(0)
		}
		execExitTest(t, "TestCmdIncompleteSSL/NoCert")
	})
	t.Run("NoKey", func(t *testing.T) {
		if os.Getenv("RUN_CRASH_TEST") == "1" {
			t.Setenv("SFILESERVER_WEBROOT", "/foo/bar")
			t.Setenv("SFILESERVER_CERT", "ssl.crt")
			parseFlags()
			// Should not reach here, ensure exit with 0 if it does
			os.Exit(0)
		}
		execExitTest(t, "TestCmdIncompleteSSL/NoKey")
	})

}
