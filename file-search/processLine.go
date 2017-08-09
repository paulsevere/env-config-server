package file_search

import (
	"regexp"
)

func processLine(query string) func(line string, lineNum int, path string) []Match {
	reg, _ := regexp.Compile(query)
	return func(line string, lineNum int, path string) []Match {
		ms := reg.FindAllStringIndex(line, 5)
		mats := make([]Match, 0)
		for _, x := range ms {
			match := Match{line: lineNum, filename: path, pos: x[0], text: line}
			mats = append(mats, match)
		}
		return mats
	}
}
