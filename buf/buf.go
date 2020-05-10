package buf

type Buffer struct {
	s string
}

func FromString(s string) *Buffer {
	return &Buffer{s: s}
}

func (b *Buffer) View() string {
	return b.s
}
