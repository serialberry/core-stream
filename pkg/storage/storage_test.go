package storage

import (
	"fmt"
	"testing"
)

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

type mockSequenceId string

func (m mockSequenceId) Next() string {
	return fmt.Sprint(m)
}

func TestGenerateSequenceImageName(t *testing.T) {
	tests := []struct {
		testCase       string
		nextId         string
		dirName        string
		baseFileName   string
		imageExtension string
		expectResult   string
	}{
		{"first sequence id", "11111", "./test", "abc", "jpeg", "test/abc-11111.jpeg"},
		{"second sequence id", "11112", "test", "abc", "jpeg", "test/abc-11112.jpeg"},
		{"third sequence id", "11113", "test/", "abc", "jpeg", "test/abc-11113.jpeg"},
		{"fourth sequence id", "11114", "parent/test", "abc", "jpeg", "parent/test/abc-11114.jpeg"},
		{"fifth sequence id", "11115", "/test", "abc", "jpeg", "/test/abc-11115.jpeg"},
	}

	for _, test := range tests {
		t.Run(test.testCase, func(assert *testing.T) {
			seq := mockSequenceId(test.nextId)
			result := GenerateSequenceImageName(seq, test.dirName, test.baseFileName, test.imageExtension)
			if test.expectResult != result {
				assert.Errorf("got [%v] want [%v]", result, test.expectResult)
			}
		})
	}
}

type mockFileProvider struct {
	canSave       bool
	generatedName string
}

func (m *mockFileProvider) Create(path string) bool {
	m.generatedName = path
	return m.canSave
}

func TestSaveImageToDisk(t *testing.T) {
	tests := []struct {
		testCase              string
		nextId                string
		dirName               string
		baseImageName         string
		canSave               bool
		expectImageNameResult string
		expectResult          bool
	}{
		{"can save frame to disk", "11111", "./test", "frame", true, "test/frame-11111.jpeg", true},
		{"cannot save frame to disk", "11112", "./test", "frame", false, "test/frame-11112.jpeg", false},
	}

	for _, test := range tests {
		t.Run(test.testCase, func(assert *testing.T) {
			fp := &mockFileProvider{canSave: test.canSave}
			seq := mockSequenceId(test.nextId)
			result := SaveImageToDisk(fp, seq, test.dirName, test.baseImageName)

			if test.expectResult != result {
				assert.Errorf("got [%v] want [%v]", result, test.expectResult)
			}

			image := fp.generatedName
			if test.expectImageNameResult != image {
				assert.Errorf("got [%v] want [%v]", image, test.expectImageNameResult)
			}
		})
	}

}
