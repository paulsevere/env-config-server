package file_search

import (
	"sync"

	"github.com/kr/fs"
)

// FileFilter
type FileFilter func(name string) bool

// WalkAndRead asds
func WalkAndRead(ff FileFilter) []Match {
	var wg sync.WaitGroup
	path, query := getArgs()
	lineHandler := processLine(query)
	walker := fs.Walk(path)
	live := true
	count := 0
	matches := make([]Match, 0)

	for live {
		live = walker.Step()
		stats := walker.Stat()
		curPath := walker.Path()
		if filterWalker(stats) {
			walker.SkipDir()
		}
		if !stats.IsDir() && stats.Mode() == 420 {
			func() {
				ch_matches := make(chan Match, 1)
				wg.Add(1)
				go readLines(curPath, ch_matches, lineHandler)
				go func() {
					defer wg.Done()
					for mat := range ch_matches {
						matches = append(matches, mat)
						count += 1
					}
				}()
			}()
		}
	}
	wg.Wait()

	return matches

}
