package ui

import "github.com/kuberlog/gt/ui/impl/mock"

func MockUi(rows int, cols int) (chan mock.Content, chan rune, Ui) {
	contentChan := make(chan mock.Content)
	keyPressChan := make(chan rune)
	ui := mock.Init(contentChan, cols, rows)
	ui.InitKeys(keyPressChan)
	return contentChan, keyPressChan, ui
}
