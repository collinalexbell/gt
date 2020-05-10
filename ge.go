package main

import (
	"kuberlog/ge/buf"
	"kuberlog/ge/win"
	"time"
)

func main() {
	b := buf.FromString("this\nis a \nbuffer")
	win := win.Init()
	win.BlitBuffer(b)
	win.Show()
	time.Sleep(time.Second * 5)
	win.Fini()
}
