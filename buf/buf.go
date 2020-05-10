package buf

import "strings"

type Marker struct {
	row, col int
}

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

func (b *Buffer) ViewLines(start int, stop int) []string {
	return b.View()[start:stop]
}

func (b *Buffer) NewMarker(row int, col int) Marker {
	m := Marker{row, col}
	return m
}

func (b *Buffer) GetLineByMarker(m Marker) string {
	line := b.ViewLines(m.row, m.row+1)
	return line[0]
}
