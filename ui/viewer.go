package ui

import (
	"github.com/kuberlog/gt/buf"
)

type Viewer struct {
	outputter Outputter
}

func InitViewer(outputter Outputter) Viewer {
	return Viewer{outputter}
}

func (viewer *Viewer) BlitBuffer(b *buf.Buffer) {
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
