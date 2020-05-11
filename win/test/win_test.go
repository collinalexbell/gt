package test

import (
	"kuberlog/ge/buf"
	"kuberlog/ge/ui/impl/mock"
	"kuberlog/ge/win"
	"testing"
)

func mockWindow() (chan mock.Content, win.Window) {
	channel := make(chan mock.Content)
	ui := mock.Init(channel, 300, 300)
	window := win.Init(ui)
	return channel, window
}

func goBlitBufferFromStr(window win.Window, str string) {
	buffer := buf.FromString(str)
	go window.BlitBuffer(buffer)
}

func testContents(expected mock.Content, actual mock.Content, t *testing.T) {
	if expected != actual {
		t.Errorf("%v != %v", expected.R, actual.R)
	}
}

func nextRow(row int, col int) (int, int) {
	row++
	col = 0
	return row, col
}

func TestBlitBufferNormal(t *testing.T) {
	channel, window := mockWindow()
	str := "this\nis a\nnew buffer\n"
	goBlitBufferFromStr(window, str)

	var actual, expected mock.Content
	row, col := 0, 0
	for _, runeVal := range str {
		if runeVal == '\n' {
			row, col = nextRow(row, col)
		} else {
			actual = <-channel
			expected = mock.Content{X: col, Y: row, R: runeVal}
			testContents(actual, expected, t)
			col++
		}
	}
}
