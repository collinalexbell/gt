package ge

import (
	"io/ioutil"
	"kuberlog/ge/buf"
	"kuberlog/ge/ui/impl/tcell"
	"kuberlog/ge/win"
	"time"
)

func open(fname string) string {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		panic("file could not be opened")
	}
	return string(b)
}

func Ge(fname string) {
	b := buf.FromString(open(fname))
	ui := tcell.Init()
	win := win.Init(ui)
	win.BlitBuffer(b)
	win.Show()
	time.Sleep(time.Second * 5)
	win.Fini()
}
