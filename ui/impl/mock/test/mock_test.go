package test

import (
	"kuberlog/ge/ui/impl/mock"
	"strings"
	"testing"
)

func TestSetContent(t *testing.T) {
	content := make(chan mock.Content)
	ui := mock.Init(content)
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
