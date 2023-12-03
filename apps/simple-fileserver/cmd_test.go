package main

import (
	"flag"
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
