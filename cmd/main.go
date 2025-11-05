package main

import (
	"os"

	"github.com/vj-2303/hello"
)

func main() {
	p := &hello.Printer{
		Output: os.Stdout,
	}
	p.Print()
}
