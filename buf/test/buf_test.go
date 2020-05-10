package test

import (
	"kuberlog/ge/buf"
	"testing"
)

func TestCreateAndView(t *testing.T) {
	str := "this\nis\na buffer\n"
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
	expected := "this\nis\na buffer\n"
	b := buf.FromString(expected)
	actual := b.ViewLines(0, 2)
	if len(actual) != 2 || "this" != actual[0] || "is" != actual[1] {
		t.Errorf("actual: %s", actual)
	}
}
