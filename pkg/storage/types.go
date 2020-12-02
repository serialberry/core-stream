package storage

type DirectoryProvider interface {
	IsExists(path string) bool // determine if named directory exists in the file system.
	Create(path string) error  // create named directory along with parent dir(s).
}

type directory struct {
	path string
}
