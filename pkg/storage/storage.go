package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gocv.io/x/gocv"
)

// Implement DirectoryProvider interface.
// Create a directory along with parent dir(s) using underline 'os' package.
// If successful returns 'nil' otherwise error.
func (s *Directory) Create(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// Implement DirectoryProvider interface.
// Checks if named directory exists using underline 'os' package.
// If directory exists, returns 'true' otherwise 'false'.
func (d *Directory) IsExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// Creates named directory along with parents if directory not found.
// Error returns only when there is a file system error.
// If directory created, return 'true' , error is nill.
// If directory not created, return 'false' , error is nil.
func CreateIfDirNotExists(provider DirectoryProvider, path string) (bool, error) {
	if provider.IsExists(path) {
		return false, nil
	}

	if err := provider.Create(path); nil != err {
		return false, err
	}

	return true, nil
}

// Implement FileProvider interface
// Save image file represent by opencv multi-dimensional array
func (f *File) Create(path string) bool {
	return gocv.IMWrite(path, f.Data)
}

// Generate unique image file name for saving to disk
// Unique name is time sequenced for sorting.
// I.E. Example file name looks like ./directory-name/base-file-name.<unix-nanoseconds>.jpeg
func GenerateSequenceImageName(seq SequenceProvider, dir string, baseName string, imageType string) string {
	return filepath.Join(dir, fmt.Sprintf("%s-%s.%s", baseName, seq.Next(), imageType))
}

// Initialise new sequence id
func NewSequenceId() SequenceId {
	return SequenceId(fmt.Sprint(time.Now().UnixNano()))
}

// Implement SequenceProvider interface
// Generate time based next id (Unix time nano seconds)
func (s SequenceId) Next() string {
	return fmt.Sprint(time.Now().UnixNano())
}

// Save opencv multi-dimensional arrary data image (frame) to disk as 'JPEG' file.
func SaveImageToDisk(provider FileProvider, seq SequenceProvider, dir string, baseImageName string) bool {
	image := GenerateSequenceImageName(seq, dir, baseImageName, "jpeg")
	return provider.Create(image)
}
