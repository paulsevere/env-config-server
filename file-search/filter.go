package file_search

import "os"

func filterWalker(stats os.FileInfo) bool {
	name := stats.Name()
	if name == "node_modules" || name == "target" || name == "dist" {
		return true
	}
	return false
}
