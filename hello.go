package hello

import (
	"fmt"
	"io"
)

type Printer struct {
	Output io.Writer
}

func (p *Printer) Print() {
	fmt.Fprintln(p.Output, "Hello, World")
}
