package main

import (
	"log"
	"net/http"
	"strconv"
)

var (
	webroot      string
	port         int
	sslCert      string
	sslKey       string
	withoutIndex bool
	debug        bool
)

func debugf(msg string, args ...any) {
	if debug {
		log.Printf("DEBUG "+msg, args...)
	}
}

type indexlessFilesystem struct {
	fs http.FileSystem
}

// Implement Open() function of interface
// Only returns without error when the path is either a file or a directory containing an index.html file
func (ifs indexlessFilesystem) Open(path string) (http.File, error) {
	f, err := ifs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		log.Printf("ERROR An unexpected error occured serving %s: %v", path, err)
		return nil, err
	}

	if s.IsDir() {
		index := path + "/index.html"
		_, err := ifs.fs.Open(index)
		if err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				debugf("Error closing file %s: %v", index, err)
				return nil, closeErr
			}
			debugf("Directory %s did not contain index.html, err=%v", path, err)
			return nil, err
		}
	}

	return f, nil
}

type indexedFilesystem struct {
	fs http.FileSystem
}

// Simple wrapper around http.Dir.Open to obtain debug information
func (ifs indexedFilesystem) Open(path string) (http.File, error) {
	debugf("Received request for %s", path)

	f, err := ifs.fs.Open(path)
	if err != nil {
		debugf("Could not open file %s: %v", path, err)
		return nil, err
	}

	return f, nil
}

// Return a filesystem with or without index enabled
func getFilesystem(path string, noIndex bool) http.FileSystem {
	fs := http.Dir(path)
	if noIndex {
		return indexlessFilesystem{fs}
	} else {
		return indexedFilesystem{fs}
	}
}

func main() {
	parseFlags()

	fs := getFilesystem(webroot, withoutIndex)
	fileServer := http.FileServer(fs)
	http.Handle("/", fileServer)

	log.Printf("Serving content from %s", webroot)

	var err error
	if sslCert == "" && sslKey == "" {
		log.Printf("Listening on :%d", port)
		err = http.ListenAndServe(":"+strconv.Itoa(port), nil)
	} else {
		log.Printf("Listening with ssl on :%d", port)
		err = http.ListenAndServeTLS(":"+strconv.Itoa(port), sslCert, sslKey, nil)
	}
	if err != nil {
		log.Fatal(err)
	}
}
