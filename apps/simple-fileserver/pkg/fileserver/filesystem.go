package fileserver

import (
	"net/http"
)

type IndexlessFilesystem struct {
	fs http.FileSystem
}

// Implement Open() function of interface
// Only returns without error when the path is either a file or a directory containing an index.html file
func (ifs IndexlessFilesystem) Open(path string) (http.File, error) {
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

type IndexedFilesystem struct {
	fs http.FileSystem
}

// Simple wrapper around http.Dir.Open to obtain debug information
func (ifs IndexedFilesystem) Open(path string) (http.File, error) {
	f, err := ifs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// Return a filesystem with or without index enabled
func CreateFilesystem(path string, index bool) http.FileSystem {
	fs := http.Dir(path)
	if index {
		return IndexedFilesystem{fs}
	} else {
		return IndexlessFilesystem{fs}
	}
}
