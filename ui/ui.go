package ui

import "github.com/kuberlog/gt/ui/event"

type Ui interface {
	SetContent(x int, y int, mainc rune)
	ScreenSize() (int, int)
	Show()
	Fini()
	PollEvent() event.Event
}
