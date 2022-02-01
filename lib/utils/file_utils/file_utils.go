package file_utils

import (
	"os"
	"path/filepath"
)

func EnsureDir(path string) error {
	return os.MkdirAll(filepath.Dir(path), os.ModeDir)
}
