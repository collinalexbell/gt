package test

import (
	"kuberlog/ge/buf"
	"kuberlog/ge/ui/impl/mock"
	"kuberlog/ge/win"
	"testing"
)

func TestBlitBufferNormal(t *testing.T) {
	channel := make(chan mock.Content)
	ui := mock.Init(channel, 300, 300)
	window := win.Init(ui)

	str := "this\nis a\nnew buffer\n"

	buffer := buf.FromString(str)
	go window.BlitBuffer(buffer)
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
