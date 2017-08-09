package file_search

import (
	"path/filepath"
)

func onlyText(name string) bool {
	ext := filepath.Ext(name)
	if ext == ".a" {
		return false
	}
	return true
}
