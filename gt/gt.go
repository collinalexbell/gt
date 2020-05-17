package gt

import (
	"fmt"
	"io/ioutil"
	"strings"

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
			if inputKey.Rune() == ':' {
				cmd := gt.cmdMode()
				if cmd == "q" {
					return
				}
			}
		}
	}
}

func (gt Gt) cmdMode() string {
	input := strings.Builder{}
	for {
		gt.viewer.DisplayCmd(input.String())
		event := gt.input.PollEvent()
		inputKey, isKey := event.(ui.InputKey)
		if isKey {
			switch inputKey.Rune() {
			case 0: //escape
				fallthrough
			case 27:
				gt.viewer.ClearCmd()
				return ""
			case 10: //enter
				fallthrough
			case 13:
				gt.viewer.ClearCmd()
				return input.String()
			default:
				input.WriteRune(inputKey.Rune())
			}
		}
	}
}
