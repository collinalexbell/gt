package main

import (
	"fmt"
	"kuberlog/ge/ge"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic(fmt.Sprintf("not enough args"))
	}
	fname := os.Args[1]
	ge.Ge(fname)
}
