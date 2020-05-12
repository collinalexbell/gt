package ge

import (
	"io/ioutil"
	"kuberlog/ge/buf"
	"kuberlog/ge/ui"
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

func Ge(fname string, ui ui.Ui) {
	b := buf.FromString(open(fname))
	win := win.Init(ui)
	win.BlitBuffer(b)
	win.Show()
	time.Sleep(time.Second * 5)
	win.Fini()
}
