package utils

import "github.com/nsf/termbox-go"

func init() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
}

func TBPrint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func Flush() {
	termbox.Flush()
}
