package main

import (
	"fmt"
	"kuberlog/ge/ge"
	"kuberlog/ge/ui/impl/tcell"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic(fmt.Sprintf("not enough args"))
	}
	fname := os.Args[1]
	ui := tcell.Init()
	ge.Ge(fname, ui)
}
