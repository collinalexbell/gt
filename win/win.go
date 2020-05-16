package win

import (
	"github.com/kuberlog/gt/buf"
	"github.com/kuberlog/gt/ui"
	"github.com/kuberlog/gt/ui/event"
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
			if y < rows && x < cols {
				window.ui.SetContent(x, y, runeValue)
			}
		}
	}
}

// todo: test this fn
func (window *Window) LookForQuit() {
	q := false
	for !q {
		e := window.ui.PollEvent()
		keyEvent, ok := e.(event.EventKey)
		if ok {
			if keyEvent.Rune() == 'q' {
				return
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
