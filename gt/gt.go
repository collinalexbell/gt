package gt

import (
	"fmt"
	"io/ioutil"

	"github.com/kuberlog/gt/buf"
	"github.com/kuberlog/gt/ui"
)

func open(fname string) string {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		panic("file could not be opened")
	}
	return string(b)
}

type Gt struct {
	viewer ui.Viewer
	input  ui.Inputter
}

func NewGt(io ui.IO) Gt {
	gt := Gt{}
	gt.InitIO(io)
	return gt
}

func (gt *Gt) InitIO(io ui.IO) {
	// do not leak the io
	// the Gt should only have access to input and viewer
	fmt.Println(io)
	gt.viewer = ui.InitViewer(io)
	gt.input = io
}

func (gt Gt) Gt(fname string) {
	b := buf.FromString(open(fname))
	gt.viewer.BlitBuffer(b)
	gt.viewer.Show()
	for {
		event := gt.input.PollEvent()
		inputKey, isKey := event.(ui.InputKey)
		if isKey {
			if inputKey.Rune() == 'q' {
				return
			}
		}
	}
}
