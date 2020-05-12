package test

import (
	"io/ioutil"
	"kuberlog/ge/ge"
	"kuberlog/ge/ui"
	"strings"
	"testing"
	"time"
)

func TestGe(t *testing.T) {
	channel, ui := ui.MockUi(300, 300)
	go ge.Ge("./ge_test.go", ui)
	b, err := ioutil.ReadFile("./ge_test.go")
	if err != nil {
		t.Errorf("could not find test file\n")
	}
	str := strings.ReplaceAll(string(b), "\n", "")
	for _, runeVal := range str {
		select {
		case content := <-channel:
			if runeVal != content.R {
				t.Errorf("%v != %v", runeVal, content.R)
			}
		case <-time.After(1 * time.Millisecond):
			t.Errorf("not enough content on channel")
		}
	}
}
