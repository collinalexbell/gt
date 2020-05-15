package event

import "time"

type Event interface {
	When() time.Time
}

type EventKey interface {
	Event
	Rune() rune
}
