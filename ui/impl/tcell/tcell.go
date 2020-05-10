package tcell

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
)

type TCellUI struct {
	screen tcell.Screen
}

func Init() *TCellUI {
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
	return &TCellUI{screen: screen}
}

func (ui *TCellUI) SetContent(x int, y int, runeVal rune) {
	ui.screen.SetContent(x, y, runeVal, []rune{}, tcell.StyleDefault)
}

func (ui *TCellUI) ScreenSize() (int, int) {
	return ui.screen.Size()
}

func (ui *TCellUI) Show() {
	ui.screen.Show()
}

func (ui *TCellUI) Fini() {
	ui.screen.Fini()
}
