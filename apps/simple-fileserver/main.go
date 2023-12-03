package main

import (
	"log"
	"net/http"
	"strconv"
)

var (
	webroot      string
	port         int
	withoutIndex bool
)

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
		return nil, err
	}

	if s.IsDir() {
		index := path + "/index.html"
		_, err := ifs.fs.Open(index)
		if err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}

// Return a filesystem with or without index enabled
func getFilesystem(path string, noIndex bool) http.FileSystem {
	fs := http.Dir(path)
	if noIndex {
		return indexlessFilesystem{fs}
	} else {
		return fs
	}
}

func main() {
	parseFlags()

	fs := getFilesystem(webroot, withoutIndex)
	fileServer := http.FileServer(fs)
	http.Handle("/", fileServer)

	log.Printf("Serving content from %s\n", webroot)

	log.Printf("Listening on :%d", port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
