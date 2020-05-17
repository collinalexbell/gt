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
