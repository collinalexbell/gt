package win

import (
	"kuberlog/ge/buf"

	"kuberlog/ge/ui"
	"kuberlog/ge/ui/impl/tcell"
)

type Window struct {
	ui ui.Ui
}

func Init() Window {
	return Window{ui: tcell.Init()}
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
