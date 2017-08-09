package file_search

import (
	"bufio"
	"os"
)

func readLines(path string, ch chan Match, cb func(line string, lineNum int, path string) []Match) {
	file, err := os.Open(path)
	count := 0
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		count += 1
		matches := cb(scanner.Text(), count, path)
		for _, x := range matches {
			ch <- x
		}

	}
	close(ch)
	return

}
