package ui

import (
	"fmt"

	"github.com/kuberlog/gt/buf"
)

type Viewer struct {
	outputter Outputter
}

func InitViewer(outputter Outputter) Viewer {
	fmt.Printf("InitViewer::outputter: %v\n", outputter)
	return Viewer{outputter}
}

func (viewer *Viewer) BlitBuffer(b *buf.Buffer) {
	fmt.Printf("BlitBuffer::viewer.ouputer: %v\n", viewer.outputter)
	cols, rows := viewer.outputter.ScreenSize()

	lines := b.View()
	for y, line := range lines {
		for x, runeValue := range line {
			if y < rows && x < cols {
				viewer.outputter.SetContent(x, y, runeValue)
			}
		}
	}
}

func (viewer *Viewer) Show() {
	viewer.outputter.Show()
}
