package test

import (
	"testing"
	"time"

	"github.com/kuberlog/gt/buf"
	"github.com/kuberlog/gt/ui"
	"github.com/kuberlog/gt/ui/impl/mock"
)

func goBlitBufferFromStr(viewer ui.Viewer, str string) {
	buffer := buf.FromString(str)
	go viewer.BlitBuffer(buffer)
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
	channel, _, viewer := MockViewer(300, 300)
	str := "this\nis a\nnew buffer\n"
	goBlitBufferFromStr(viewer, str)

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
	channel, _, viewer := MockViewer(3, cols)
	str := "abc\ndef\n"
	goBlitBufferFromStr(viewer, str)
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
	channel, _, viewer := MockViewer(rows, 2)
	str := "ab\ncd\nef\n"
	goBlitBufferFromStr(viewer, str)

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
