package win

import (
	"fmt"
	"kuberlog/ge/buf"
	"os"

	"github.com/gdamore/tcell"
)

type Window struct {
	screen tcell.Screen
}

func Init() Window {
	screen, errors := tcell.NewScreen()
	if errors != nil {
		fmt.Fprintf(os.Stderr, "Failed to create screen")
		os.Exit(1)
	}
	if e := screen.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize screen")
		os.Exit(1)
	}

	screen.SetStyle(tcell.StyleDefault)
	screen.Clear()
	return Window{screen: screen}
}

func (window *Window) BlitBuffer(b *buf.Buffer) {
	cols, rows := window.screen.Size()
	lines := b.View()
	for y, line := range lines {
		for x, runeValue := range line {
			if y <= rows && x <= cols {
				window.screen.SetContent(x, y, runeValue, []rune{}, tcell.StyleDefault)
			}
		}
	}
}

func (window *Window) Show() {
	window.screen.Show()
}

func (window *Window) Fini() {
	window.screen.Fini()
}
