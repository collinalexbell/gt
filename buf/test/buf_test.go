package test

import (
	"kuberlog/ge/buf"
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
	marker := b.NewMarker(1, 0)
	line := b.GetLineByMarker(marker)
	if line != "is" {
		t.Errorf("%s != %s", "is", line)
	}
}
