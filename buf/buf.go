package buf

import "strings"

type Marker struct {
	row, col int
}

type Buffer struct {
	s       []string
	markers []Marker
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

func (b *Buffer) Mark(row int, col int) Marker {
	m := Marker{row, col}
	b.markers = append(b.markers, m)
	return m
}

func (b *Buffer) DeleteMarker(marker Marker) {
	for index := range b.markers {
		if marker == b.markers[index] {
			b.markers = append(b.markers[0:index], b.markers[index+1:]...)
			return // Only want to delete the first marker? I really should do this with pointers.
		}
	}
}

func (b *Buffer) GetMarkers() []Marker {
	return b.markers
}

func (b *Buffer) GetLineByMarker(m Marker) string {
	line := b.ViewLines(m.row, m.row+1)
	return line[0]
}

func (b *Buffer) DeleteLineByMarker(m Marker) {
	if m.row >= len(b.s) {
		return
	}
	if m.row == len(b.s)-1 {
		b.s = b.s[0:m.row]
	} else {
		b.s = append(b.s[:m.row], b.s[m.row+1:]...)
	}

}
