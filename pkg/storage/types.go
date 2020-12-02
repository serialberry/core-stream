package storage

type DirectoryProvider interface {
	IsDirectoryExists(path string) bool // determine if named directory exists in the file system.
	CreateDirectory(path string) error  // create named directory along with parent dir(s).
}

type writer struct {
	dir string
}
