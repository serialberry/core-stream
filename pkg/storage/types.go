package storage

import "gocv.io/x/gocv"

type DirectoryProvider interface {
	IsExists(path string) bool // determine if named directory exists in the file system.
	Create(path string) error  // create named directory along with parent dir(s).
}

type Directory struct {
	Path string
}

type FileProvider interface {
	Create(path string) bool
}

type File struct {
	Data gocv.Mat
	Path string
}

type SequenceProvider interface {
	Next() string
}

type SequenceId string
