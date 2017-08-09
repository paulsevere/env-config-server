package file_search

import "os"

func getArgs() (string, string) {
	args := os.Args[1:]
	query := args[1]
	path := args[0]
	return path, query
}
