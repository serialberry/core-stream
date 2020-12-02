package storage

import (
	"os"
)

// Implement FileSystemWriter interface.
// Create a directory along with parent dir(s) using underline 'os' package.
// If successful returns 'nil' otherwise error.
func (s *directory) CreateDirectory(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// Implement FileSystemWriter interface.
// Checks if named directory exists using underline 'os' package.
// If directory exists, returns 'true' otherwise 'false'.
func (d *directory) IsDirectoryExists(path string) bool {
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
