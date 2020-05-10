package main

import (
	"kuberlog/ge/buf"
	"kuberlog/ge/ui"
	"time"
)

func main() {
	b := buf.FromString("this\nis a \nbuffer")
	ui := ui.Init()
	ui.BlitBuffer(b)
	ui.Show()
	time.Sleep(time.Second * 5)
	ui.Fini()
}
