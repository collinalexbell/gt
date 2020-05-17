package ui

import "time"

type Inputter interface {
	PollEvent() InputEvent
}

type InputEvent interface {
	When() time.Time
}

type InputKey interface {
	InputEvent
	Rune() rune
}

type Outputter interface {
	SetContent(x int, y int, mainc rune)
	ScreenSize() (int, int) // cols, rows
	Show()
	ShowCursor(col int, row int)
	HideCursor()
}

type IO interface {
	Outputter
	Inputter
}
