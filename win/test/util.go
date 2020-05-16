package test

import (
	"github.com/kuberlog/gt/ui"
	"github.com/kuberlog/gt/ui/impl/mock"
	"github.com/kuberlog/gt/win"
)

func MockWindow(rows int, cols int) (chan mock.Content, chan rune, win.Window) {
	contentChan, keyPressChan, ui := ui.MockUi(rows, cols)
	window := win.Init(ui)
	return contentChan, keyPressChan, window
}
