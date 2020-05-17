package ui

type Outputter interface {
	SetContent(x int, y int, mainc rune)
	ScreenSize() (int, int)
	Show()
}
