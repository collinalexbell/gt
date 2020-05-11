package win

import (
	"kuberlog/ge/buf"

	"kuberlog/ge/ui"
)

type Window struct {
	ui ui.Ui
}

func Init(ui ui.Ui) Window {
	return Window{ui: ui}
}

func (window *Window) BlitBuffer(b *buf.Buffer) {
	cols, rows := window.ui.ScreenSize()
	lines := b.View()
	for y, line := range lines {
		for x, runeValue := range line {
			if y <= rows && x <= cols {
				window.ui.SetContent(x, y, runeValue)
			}
		}
	}
}

func (window *Window) Show() {
	window.ui.Show()
}

func (window *Window) Fini() {
	window.ui.Fini()
}
