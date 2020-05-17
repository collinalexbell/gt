package test

import (
	"io/ioutil"
	"strings"
	"testing"
	"time"

	gtpack "github.com/kuberlog/gt/gt"
	ui "github.com/kuberlog/gt/ui/test"
)

func TestGt(t *testing.T) {
	channel, keyChan, ui := ui.MockUi(300, 300)
	gt := gtpack.NewGt(ui)
	go gt.Gt("./gt_test.go")
	b, err := ioutil.ReadFile("./gt_test.go")
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
	keyChan <- 'q'
}
