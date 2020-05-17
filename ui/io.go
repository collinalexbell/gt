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
	ScreenSize() (int, int)
	Show()
}

type IO interface {
	Outputter
	Inputter
}
