package test

import (
	"kuberlog/ge/buf"
	"kuberlog/ge/ui/impl/mock"
	"kuberlog/ge/win"
	"testing"
	"time"
)

func mockWindow(rows int, cols int) (chan mock.Content, win.Window) {
	channel := make(chan mock.Content)
	ui := mock.Init(channel, cols, rows)
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

func TestBlitBuffer_Normal(t *testing.T) {
	channel, window := mockWindow(300, 300)
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

func TestBlitBuffer_TooWide(t *testing.T) {
	cols := 2
	channel, window := mockWindow(3, cols)
	str := "abc\ndef\n"
	goBlitBufferFromStr(window, str)
	var actual, expected mock.Content
	row, col := 0, 0
	for _, runeVal := range str {
		if runeVal == '\n' {
			row, col = nextRow(row, col)
		} else {
			if col < cols {
				actual = <-channel
				expected = mock.Content{X: col, Y: row, R: runeVal}
				testContents(actual, expected, t)
				col++
			}
		}
	}
}

func TestBlitBuffer_TooLong(t *testing.T) {
	rows := 2
	channel, window := mockWindow(rows, 2)
	str := "ab\ncd\nef\n"
	goBlitBufferFromStr(window, str)

	var actual, expected mock.Content
	row, col := 0, 0
	for _, runeVal := range str {
		if runeVal == '\n' {
			row, col = nextRow(row, col)
		} else {
			if row < rows {
				actual = <-channel
				expected = mock.Content{X: col, Y: row, R: runeVal}
				testContents(actual, expected, t)
				col++
			}
		}
	}
	//Wait for any extra row data to get on the channel
	time.Sleep(50 * time.Millisecond)
	select {
	case <-channel:
		t.Errorf("channel should be exhauseted, but another row was added to it")
	default:
	}
}
