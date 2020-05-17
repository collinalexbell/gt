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

func (viewer *Viewer) Debug(s string) {
	cols, rows := viewer.outputter.ScreenSize()
	colInd := cols - len(s) - 1
	row := rows - 1

	for _, r := range s {
		viewer.outputter.SetContent(colInd, row, r)
		colInd++
	}
	viewer.outputter.Show()
}

func (viewer *Viewer) DisplayCmd(cmd string) {
	_, rows := viewer.outputter.ScreenSize()
	row := rows - 1
	colInd := 1

	viewer.outputter.SetContent(colInd-1, row, ':')
	for _, r := range cmd {
		viewer.outputter.SetContent(colInd, row, r)
		colInd++
	}
	viewer.outputter.Show()
}

func (viewer *Viewer) ClearCmd() {
	cols, rows := viewer.outputter.ScreenSize()
	row := rows - 1
	for col := 0; col < cols; col++ {
		viewer.outputter.SetContent(col, row, ' ')
	}
	viewer.outputter.Show()
}
