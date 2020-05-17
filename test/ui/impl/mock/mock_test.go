package mock

import (
	"strings"
	"testing"

	"github.com/kuberlog/gt/ui"
	"github.com/kuberlog/gt/ui/impl/mock"
)

func TestSetContent(t *testing.T) {
	content := make(chan mock.Content)
	var ui *mock.MockUi = mock.InitMockUi(content, 100, 100)
	str := "this\nis\na buffer"
	for row, line := range strings.Split(str, "\n") {
		for col, r := range line {
			go ui.SetContent(col, row, r)
			actual := <-content
			expected := mock.Content{X: col, Y: row, R: r}
			if actual != expected {
				t.Errorf("actual: %v != expected: %v", actual, expected)
			}
		}
	}
}

func TestDefaultScreenSize(t *testing.T) {
	content := make(chan mock.Content)
	ui := mock.InitMockUi(content, 140, 100)
	x, y := ui.ScreenSize()
	if x != 140 || y != 100 {
		t.Fail()
	}
}

func TestPollKeyEvent(t *testing.T) {
	content := make(chan mock.Content)
	mockUi := mock.InitMockUi(content, 140, 100)
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
