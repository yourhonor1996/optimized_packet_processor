package processor

import (
	"path/filepath"
)

func GenerateResultFileName(name string) string {
	directory := filepath.Dir(name)
	return filepath.Join(directory, "result_"+filepath.Base(name))
}
