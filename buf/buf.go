package buf

import "strings"

type Buffer struct {
	s []string
}

func FromString(s string) *Buffer {
	slice := strings.Split(s, "\n")
	return &Buffer{s: slice}
}

func (b *Buffer) View() []string {
	return b.s
}

func (b *Buffer) ViewLines(start int64, stop int64) []string {
	return b.View()[start:stop]
}
