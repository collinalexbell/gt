package mock

type Content struct {
	X, Y int
	R    rune
}

type MockUi struct {
	content chan Content
}

func Init(content chan Content) *MockUi {
	return &MockUi{content}
}

func (ui *MockUi) SetContent(x int, y int, runeVal rune) {
	ui.content <- Content{x, y, runeVal}
}
