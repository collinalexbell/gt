package ui

import "github.com/kuberlog/gt/ui/impl/mock"

func MockUi(rows int, cols int) (chan mock.Content, Ui) {
	channel := make(chan mock.Content)
	ui := mock.Init(channel, cols, rows)
	return channel, ui
}
