package ui

import (
	"fmt"
	"kuberlog/ge/buf"
	"os"

	"github.com/gdamore/tcell"
)

type Ui struct {
	screen tcell.Screen
}

func Init() Ui {
	screen, errors := tcell.NewScreen()
	if errors != nil {
		fmt.Fprintf(os.Stderr, "Failed to create screen")
		os.Exit(1)
	}
	if e := screen.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize screen")
		os.Exit(1)
	}

	screen.SetStyle(tcell.StyleDefault)
	screen.Clear()
	return Ui{screen: screen}
}

func (ui *Ui) BlitBuffer(b *buf.Buffer) {
	cols, rows := ui.screen.Size()
	lines := b.View()
	for y, line := range lines {
		for x, runeValue := range line {
			if y <= rows && x <= cols {
				ui.screen.SetContent(x, y, runeValue, []rune{}, tcell.StyleDefault)
			}
		}
	}
}

func (ui *Ui) Show() {
	ui.screen.Show()
}

func (ui *Ui) Fini() {
	ui.screen.Fini()
}
