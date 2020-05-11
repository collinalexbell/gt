package main

import (
	"kuberlog/ge/buf"
	"kuberlog/ge/ui/impl/tcell"
	"kuberlog/ge/win"
	"time"
)

func main() {
	b := buf.FromString("this\nis a \nbuffer")
	ui := tcell.Init()
	win := win.Init(ui)
	win.BlitBuffer(b)
	win.Show()
	time.Sleep(time.Second * 5)
	win.Fini()
}
