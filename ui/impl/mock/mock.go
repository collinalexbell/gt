package mock

import (
	"time"

	"github.com/kuberlog/gt/ui/event"
)

type Event struct{}

func (Event) When() time.Time {
	return time.Now()
}

type Content struct {
	X, Y int
	R    rune
}

type MockUi struct {
	content      chan Content
	sizeX, sizeY int
}

func Init(content chan Content, sizeX int, sizeY int) *MockUi {
	return &MockUi{content, sizeX, sizeY}
}

func (ui *MockUi) SetContent(x int, y int, runeVal rune) {
	ui.content <- Content{x, y, runeVal}
}

func (ui *MockUi) ScreenSize() (int, int) {
	return ui.sizeX, ui.sizeY
}

func (ui *MockUi) Show() {}

func (ui *MockUi) Fini() {}

func (ui *MockUi) PollEvent() event.Event {
	return Event{}
}
