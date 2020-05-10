package test

import (
	"kuberlog/ge/buf"
	"testing"
)

func TestCreateAndView(t *testing.T) {
	expected := "this\nis\na buffer\n"
	b := buf.FromString(expected)
	actual := b.View()
	if expected != actual {
		t.Fail()
	}
}
