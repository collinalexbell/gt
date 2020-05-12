package test

import (
	"kuberlog/ge/ui"
	"kuberlog/ge/ui/impl/mock"
	"kuberlog/ge/win"
)

func MockWindow(rows int, cols int) (chan mock.Content, win.Window) {
	channel, ui := ui.MockUi(rows, cols)
	window := win.Init(ui)
	return channel, window
}
