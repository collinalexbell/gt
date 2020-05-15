package main

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/kuberlog/gt/ui/impl/tcell"

	"github.com/kuberlog/gt/gt"
)

func redirectStderr(f *os.File) {
	err := syscall.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		log.Fatalf("Failed to redirect stderr to file: %v", err)
	}
}

func main() {
	if len(os.Args) < 2 {
		panic(fmt.Sprintf("not enough args"))
	}
	if len(os.Args) > 2 {
		errFile, err := os.Create(os.Args[2])
		if err != nil {
			panic("cant redirect err")
		}
		redirectStderr(errFile)
	}
	fname := os.Args[1]

	ui := tcell.Init()
	gt.Gt(fname, ui)
}
