package test

import (
	"github.com/kuberlog/gt/ui"
	"github.com/kuberlog/gt/ui/impl/mock"
	"strings"
	"testing"
)

func TestSetContent(t *testing.T) {
	content := make(chan mock.Content)
	var ui ui.Ui = mock.Init(content, 100, 100)
	str := "this\nis\na buffer"
	for row, line := range strings.Split(str, "\n") {
		for col, r := range line {
			go ui.SetContent(col, row, r)
			actual := <-content
			expected := mock.Content{col, row, r}
			if actual != expected {
				t.Errorf("actual: %v != expected: %v", actual, expected)
			}
		}
	}
}

func TestDefaultScreenSize(t *testing.T) {
	content := make(chan mock.Content)
	var ui ui.Ui = mock.Init(content, 140, 100)
	x, y := ui.ScreenSize()
	if x != 140 || y != 100 {
		t.Fail()
	}
}
