package mock

import (
	"time"
)

type EventKey struct {
	r rune
}

func (EventKey) When() time.Time {
	return time.Now()
}

func (e EventKey) Rune() rune {
	return e.r
}

type Content struct {
	X, Y int
	R    rune
}

type MockUi struct {
	content      chan Content
	sizeX, sizeY int
	keys         chan rune
}

func InitMockUi(content chan Content, sizeX int, sizeY int) *MockUi {
	return &MockUi{content, sizeX, sizeY, nil}
}

func (ui *MockUi) InitKeys(keys chan rune) {
	ui.keys = keys
}

func (ui *MockUi) SetContent(x int, y int, runeVal rune) {
	ui.content <- Content{x, y, runeVal}
}

func (ui *MockUi) ScreenSize() (int, int) {
	return ui.sizeX, ui.sizeY
}

func (ui *MockUi) Show() {}

func (ui *MockUi) Fini() {}

func (ui *MockUi) PollEvent() EventKey {
	if ui.keys == nil {
		// todo: return an err
	}
	key := <-ui.keys
	return EventKey{key}
}
