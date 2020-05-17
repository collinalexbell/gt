package test

import (
	"github.com/kuberlog/gt/ui"
	"github.com/kuberlog/gt/ui/impl/mock"
)

func MockUi(rows int, cols int) (chan mock.Content, chan rune, *mock.MockUi) {
	contentChan := make(chan mock.Content)
	keyPressChan := make(chan rune)
	mock := mock.InitMockUi(contentChan, cols, rows)
	mock.InitKeys(keyPressChan)
	return contentChan, keyPressChan, mock
}

func MockViewer(rows int, cols int) (chan mock.Content, chan rune, ui.Viewer) {
	contentChan, keyPressChan, mockUi := MockUi(rows, cols)
	viewer := ui.InitViewer(mockUi)
	return contentChan, keyPressChan, viewer
}
