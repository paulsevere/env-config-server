package file_search

import "fmt"

type Match struct {
	line     int
	pos      int
	filename string
	text     string
}

func main() {
	fmt.Printf("Found %v matches", len(WalkAndRead(onlyText)))

}

func waste(interface{}) {

}
