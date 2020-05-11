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

func blitBufferFromStr(window win.Window, str string) {
	buffer := buf.FromString(str)
	go window.BlitBuffer(buffer)
}

func TestBlitBufferNormal(t *testing.T) {
	channel, window := mockWindow()

	str := "this\nis a\nnew buffer\n"
	blitBufferFromStr(window, str)

	row, col := 0, 0
	for _, r := range str {
		if r == '\n' {
			row++
			col = 0
		} else {
			c := <-channel
			if c.R != r {
				t.Errorf("%v != %v", c.R, r)
			}
			col++
		}
	}
}
