package test

import (
	"kuberlog/ge/buf"
	"strings"
	"testing"
)

var str = "this\nis\na buffer\n"

func TestCreateAndView(t *testing.T) {
	b := buf.FromString(str)
	actual := b.View()
	expected := []string{"this", "is", "a buffer"}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Fail()
		}
	}
}

func TestViewWithIndices(t *testing.T) {
	b := buf.FromString(str)
	actual := b.ViewLines(0, 2)
	if len(actual) != 2 || "this" != actual[0] || "is" != actual[1] {
		t.Errorf("actual: %s", actual)
	}
}

func TestGetLineByMarker(t *testing.T) {
	b := buf.FromString(str)
	marker := b.Mark(1, 0)
	line := b.GetLineByMarker(marker)
	if line != "is" {
		t.Errorf("%s != %s", "is", line)
	}
}

func TestDeleteLineByMarker_MiddleLine(t *testing.T) {
	b := buf.FromString(str)
	marker := b.Mark(1, 0)

	b.DeleteLineByMarker(marker)
	lines := b.View()
	strSlice := strings.Split(str, "\n")
	expected := append(strSlice[:1], strSlice[2:]...)
	for i, line := range lines {
		if line != expected[i] {
			t.Errorf("%s != %s", line, expected[i])
		}
	}
}

func TestDeleteLineByMarker_MarkerRowToBig(t *testing.T) {
	b := buf.FromString(str)
	marker := b.Mark(10, 0)

	b.DeleteLineByMarker(marker)
	lines := b.View()

	// Should do nothing
	strSlice := strings.Split(str, "\n")
	for i, line := range lines {
		if line != strSlice[i] {
			t.Errorf("%s != %s", line, strSlice[i])
		}
	}
}

func TestDeleteLineByMarker_MarkerRowIsZero(t *testing.T) {
	b := buf.FromString(str)
	marker := b.Mark(0, 0)

	b.DeleteLineByMarker(marker)
	lines := b.View()

	// Should do nothing
	strSlice := strings.Split(str, "\n")
	expected := strSlice[1:]
	for i, line := range lines {
		if line != expected[i] {
			t.Errorf("%s != %s", line, expected[i])
		}
	}
}
