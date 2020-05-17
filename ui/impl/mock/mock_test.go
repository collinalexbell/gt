package mock

import (
	"strings"
	"testing"

	"github.com/kuberlog/gt/ui"
)

func TestSetContent(t *testing.T) {
	content := make(chan Content)
	var ui *MockUi = InitMockUi(content, 100, 100)
	str := "this\nis\na buffer"
	for row, line := range strings.Split(str, "\n") {
		for col, r := range line {
			go ui.SetContent(col, row, r)
			actual := <-content
			expected := Content{X: col, Y: row, R: r}
			if actual != expected {
				t.Errorf("actual: %v != expected: %v", actual, expected)
			}
		}
	}
}

func TestDefaultScreenSize(t *testing.T) {
	content := make(chan Content)
	ui := InitMockUi(content, 140, 100)
	x, y := ui.ScreenSize()
	if x != 140 || y != 100 {
		t.Fail()
	}
}

func TestPollKeyEvent(t *testing.T) {
	content := make(chan Content)
	mockUi := InitMockUi(content, 140, 100)
	keys := make(chan rune)
	mockUi.InitKeys(keys)
	go func() { keys <- 'q' }()
	event := mockUi.PollEvent()
	eventKey, ok := event.(ui.InputKey)
	if !ok {
		t.Errorf("event is not a EventKey")
	}
	r := eventKey.Rune()
	if r != 'q' {
		t.Errorf("event Key did not contain the correct rune")
	}
}