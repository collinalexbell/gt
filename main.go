package main

import (
	"fmt"
	"os"

	"github.com/kuberlog/gt/ui/impl/tcell"

	"github.com/kuberlog/gt/gt"
)

func main() {
	if len(os.Args) < 2 {
		panic(fmt.Sprintf("not enough args"))
	}
	fname := os.Args[1]
	ui := tcell.Init()
	gt.Gt(fname, ui)
}
