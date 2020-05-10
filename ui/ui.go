package ui

type Ui interface {
	SetContent(x int, y int, mainc rune)
	ScreenSize() (int, int)
	Show()
	Fini()
}
