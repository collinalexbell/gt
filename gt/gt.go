package gt

import (
	"io/ioutil"
	"time"

	"github.com/kuberlog/gt/buf"
	"github.com/kuberlog/gt/ui"
	"github.com/kuberlog/gt/win"
)

func open(fname string) string {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		panic("file could not be opened")
	}
	return string(b)
}

func Gt(fname string, ui ui.Ui) {
	b := buf.FromString(open(fname))
	win := win.Init(ui)
	win.BlitBuffer(b)
	win.Show()
	time.Sleep(time.Second * 5)
	win.Fini()
}
