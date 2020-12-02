package storage

import "testing"

type mockFileSystemError struct {
	message string
}

type mockDirectoryProvider struct {
	existingDir bool
	err         error
}

// Mock error interface
func (m *mockFileSystemError) Error() string {
	return m.message
}

// Mock DirectoryProvider interface
func (m *mockDirectoryProvider) Create(name string) error {
	return m.err
}

// Mock DirectoryProvider interface
func (m *mockDirectoryProvider) IsExists(path string) bool {
	return m.existingDir
}

func TestCreateIfDirNotExists(t *testing.T) {
	tests := []struct {
		testCase      string
		ifExistingDir bool
		expectError   error
		expectResult  bool
	}{
		{"create new directory", false, nil, true},
		{"skip creating directory", true, nil, false},
		{"skip creating directory when error raised", false, &mockFileSystemError{message: "filesystem error"}, false},
	}

	for _, test := range tests {
		t.Run(test.testCase, func(assert *testing.T) {
			fs := &mockDirectoryProvider{existingDir: test.ifExistingDir, err: test.expectError}
			result, err := CreateIfDirNotExists(fs, "./test")
			if err != test.expectError {
				assert.Errorf("got [%v] want [%v]", err, test.expectError)
			}
			if result != test.expectResult {
				assert.Errorf("got [%v] want [%v]", result, test.expectResult)
			}
		})
	}
}
