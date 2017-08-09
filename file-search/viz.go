package file_search

import "github.com/buger/goterm"

// WaitDots asd
type WaitDots struct {
	len  int
	curr int
	cha  string
}

// Render asa
func (w *WaitDots) Render() {
	if w.curr < w.len {
		w.curr++
	} else {
		w.curr = 0
		goterm.Clear()
		goterm.MoveCursor(1, 1)
	}
	goterm.Println(w.cha)
}
