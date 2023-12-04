package main

import (
	"net/http"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testMatrixGetFilesystem struct {
	name         string
	path         string
	noIndex      bool
	expectedType string
}

func TestGetFilesystem(t *testing.T) {

	testMatrix := []testMatrixGetFilesystem{
		{
			name:         "withIndex",
			path:         "foo",
			noIndex:      false,
			expectedType: "main.indexedFilesystem",
		},
		{
			name:         "withoutIndex",
			path:         "foo",
			noIndex:      true,
			expectedType: "main.indexlessFilesystem",
		},
	}

	for _, tCase := range testMatrix {
		t.Run(tCase.name, func(t *testing.T) {
			fs := getFilesystem(tCase.path, tCase.noIndex)

			assert.Equal(t, tCase.expectedType, reflect.TypeOf(fs).String())
		})
	}
}

func TestIndexlessFilesystem(t *testing.T) {
	fs := indexlessFilesystem{http.Dir("./testdata")}

	t.Run("DirWithoutIndexFile", func(t *testing.T) {
		assert := assert.New(t)

		f, err := fs.Open("/")
		assert.Nil(f)
		if assert.Error(err) {
			switch osString := strings.ToLower(runtime.GOOS); osString {
			case "windows":
				assert.ErrorContains(err, "The system cannot find the file specified")
			case "linux":
				assert.ErrorContains(err, "open testdata/index.html: no such file or directory")
			default:
				t.Fatalf("Unknown OS %s", osString)
			}
		}
	})

	testMatrix := map[string]string{
		"File":             "/test.html",
		"DirWithIndexFile": "/testdir",
	}

	for name, path := range testMatrix {
		t.Run(name, func(t *testing.T) {
			f, err := fs.Open(path)
			if assert.Nil(t, err) {
				err = f.Close()
				if err != nil {
					t.Fatalf("Unexpected error closing file: %v", err)
				}
			}
		})
	}
}
